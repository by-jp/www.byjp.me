package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/by-jp/www.byjp.me/tools/shared"

	synd "github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

const webmentions = "https://webmention.io/api/mentions.jf2?domain=%s&token=%s"
const siteTarget = "https://www.byjp.me/"

type Config struct {
	Domain string
	Token  string
	Root   string
}

func main() {
	var c Config
	check(shared.LoadConfig("webmentionio", &c))

	mentions, err := retrieveMentions(c)
	check(err)

	pathChecker := checkValidPath(path.Join(c.Root, "content"))
	for _, m := range mentions.Children {
		pathStr, interaction, err := parseJF2(m, pathChecker)
		switch err {
		case nil:
			// Proceed
		case ErrNoEntry, ErrIsPrivate, ErrIncorrectTarget:
			// Skip
			continue
		default:
			// Stop
			check(err)
		}

		dst := path.Join(c.Root, "data/interactions", pathStr+".json")
		addErr := addInteraction(dst, interaction)
		if addErr != nil {
			fmt.Fprintf(os.Stderr, "Couldn't save interaction %s to %s: %v\n", interaction.GUID, dst, addErr)
		}
	}
}

func checkValidPath(contentDir string) func(string) bool {
	return func(pathStr string) bool {
		pathDir := path.Join(contentDir, pathStr)
		std, err := os.Stat(pathDir)
		if err == nil && std.IsDir() {
			return true
		}

		pathFile := pathDir + ".md"
		stf, err := os.Stat(pathFile)
		if err == nil && stf.Mode().IsRegular() {
			return true
		}

		return false
	}
}

func addInteraction(jsonPath string, newIn synd.Interaction) error {
	inf := InteractionFile{}

	rf, err := os.Open(jsonPath)
	if err == nil {
		if err := json.NewDecoder(rf).Decode(&inf); err != nil && err != io.EOF {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	rf.Close()

	added := false
	for idx, in := range inf.Interactions {
		if in.GUID == newIn.GUID {
			inf.Interactions[idx] = newIn
			added = true
		}
	}
	if !added {
		inf.Interactions = append(inf.Interactions, newIn)
	}

	if err := os.MkdirAll(path.Dir(jsonPath), 0755); err != nil {
		return err
	}

	wf, err := os.OpenFile(jsonPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer wf.Close()

	return json.NewEncoder(wf).Encode(inf)
}

var ErrNoEntry = errors.New("JF2 item isn't an entry")
var ErrIsPrivate = errors.New("JF2 item is set to private")
var ErrIncorrectTarget = errors.New("JF2 item describes a different target site")

func parseJF2(m Mention, isValidPath func(string) bool) (string, synd.Interaction, error) {
	if m.Type != "entry" {
		return "", synd.Interaction{}, ErrNoEntry
	}
	if m.WebmentionPrivate {
		return "", synd.Interaction{}, ErrIsPrivate
	}
	if !strings.HasPrefix(m.WebmentionTarget, siteTarget) {
		return "", synd.Interaction{}, ErrIncorrectTarget
	}
	if strings.HasPrefix(m.URL, siteTarget) {
		return "", synd.Interaction{}, ErrIncorrectTarget
	}

	sitePath := strings.TrimRight(m.WebmentionTarget[len(siteTarget):], "/")
	if !isValidPath(sitePath) {
		return "", synd.Interaction{}, ErrIncorrectTarget
	}

	i := synd.Interaction{
		GUID:    fmt.Sprintf("webmentions.io#%d", m.WebmentionID),
		URL:     m.URL,
		Comment: m.Content.Text,
		Author: synd.Author{
			// TODO: Handle Bluesky no name on likes
			Name: m.Author.Name,
			URL:  m.Author.URL,
			// TODO: Photo?
		},
	}

	p, err := time.Parse(time.RFC3339, m.Published)
	if err != nil || p.IsZero() {
		p, _ = time.Parse(time.RFC3339, m.WebmentionReceived)
	}
	i.Timestamp = p

	// TODO: very long content?
	// TODO: If wm-source != URL then acting on a reference, not on a syndicated post

	switch m.WebmentionProperty {
	case "like-of":
		if strings.Contains(m.Author.URL, "/@") {
			i.Emoji = "⭐️"
		} else {
			i.Emoji = "♥️"
		}

	case "repost-of":
		i.Emoji = "🔁"
	case "in-reply-to":
		i.Emoji = "💬"
	case "mention-of":
		i.Comment = "_A wibble_"
		i.Emoji = "🗣️"
	}

	return sitePath, i, nil
}

func retrieveMentions(c Config) (Mentions, error) {
	url := webmentionsURL(c.Domain, c.Token)
	res, err := http.Get(url)
	if err != nil {
		return Mentions{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return Mentions{}, fmt.Errorf("Non 200 response code: %v\n%s", res.StatusCode, body)
	}

	var m Mentions
	return m, json.NewDecoder(res.Body).Decode(&m)
}

func webmentionsURL(domain, token string) string {
	return fmt.Sprintf(webmentions, domain, token)
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
