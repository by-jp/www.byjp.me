package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
)

var tagMap = map[string]string{
	"helloiamrob":   "robhunt",
	"chrismhs":      "chris",
	"jphastings":    "jp",
	"lennym":        "lenny",
	"edds":          "edds",
	"hungoverdrawn": "hungoverdrawn",
}

func check(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n  %v\n", msg, err)
		os.Exit(1)
	}
}

type closer interface {
	Close() error
}

func doClose(c closer, msg string) {
	check(c.Close(), msg)
}

type frontMatter struct {
	Date      string
	Tags      []string
	Location  fmLoc  `yaml:"location,omitempty"`
	InReplyTo string `yaml:"in_reply_to,omitempty"`
}

type fmLoc struct {
	Name      string `yaml:"name,omitempty"`
	Latitude  float64
	Longitude float64
}

type tweet struct {
	Tweet struct {
		ID                string `json:"id_str"`
		FullText          string `json:"full_text"`
		CreatedAt         string `json:"created_at"`
		CreationTimestamp int64  `json:"creation_timestamp"`
		InReplyTo         string `json:"in_reply_to_user_id_str"`
		InReplyToStatus   string `json:"in_reply_to_status_id_str"`
		Coordinates       struct {
			LonLat []string `json:"coordinates"`
			Type   string   `json:"type"`
		} `json:"coordinates"`
		Entities struct {
			Media []struct {
				ID          string `json:"id_str"`
				MediaURL    string `json:"media_url"`
				ExpandedURL string `json:"expanded_url"`
			}
		} `json:"entities"`
	} `json:"tweet"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path/to/twitter/archive.zip> <path/to/hugo>\n", os.Args[0])
	}

	archive := os.Args[1]
	hugo := os.Args[2]
	outputDir := path.Join(hugo, "content", "notes", "twitter")

	check(os.MkdirAll(outputDir, 0755), "Unable to create twitter directory in notes")

	zf, err := zip.OpenReader(archive)
	check(err, "Unable to open twitter archive")
	defer doClose(zf, "Unable to close zipfile")

	notesCount, mediaMap, err := createNotes(zf, outputDir)
	check(err, "Unable to create hugo posts for your twitter data")
	check(copyMedia(zf, mediaMap), "Unable to copy media to your hugo blog")

	fmt.Printf("Success! %d Twitter posts (with %d images and videos) were added to your hugo blog.\n", notesCount, len(mediaMap))
}

var notesFile = "data/tweets.js"

func createNotes(zf *zip.ReadCloser, outputDir string) (int, map[string]string, error) {
	for _, f := range zf.File {
		if notesFile != f.Name {
			continue
		}

		jf, err := f.Open()
		if err != nil {
			return 0, nil, err
		}
		defer doClose(jf, "Unable to close posts file within archive")

		return notesFromFile(jf, outputDir)
	}

	return 0, nil, fmt.Errorf("no %s file found in zip file", notesFile)
}

func notesFromFile(r io.Reader, outputDir string) (int, map[string]string, error) {
	if err := readUntil(r, " = ", 1024); err != nil {
		return 0, nil, err
	}

	notesCount := 0
	mediaMap := make(map[string]string)
	dec := json.NewDecoder(r)

	// Opening [
	tok, err := dec.Token()
	if err != nil {
		return 0, nil, err
	}
	if fmt.Sprintf("%s", tok) != "[" {
		return 0, nil, fmt.Errorf("tweets start with %s instead of '['", tok)
	}

	for dec.More() {
		var t tweet
		if err := dec.Decode(&t); err != nil {
			return notesCount, mediaMap, fmt.Errorf("unable to decode JSON: %w", err)
		}

		// TODO: Collect self user ID from archive somehow
		if err := tweetToNote(t, mediaMap, outputDir, "12721"); err != nil {
			return notesCount, mediaMap, err
		}
		notesCount++
	}

	return notesCount, mediaMap, nil
}

var hashtag = regexp.MustCompile(`[\s\[\(,…][#@][a-zA-Z0-9_]{3,}\b`)
var urlShortener = regexp.MustCompile(`https?://(t.co|j.mp|bit.ly|tinyurl.com)/\w+`)
var twitpic = regexp.MustCompile(`https?://twitpic.com/[a-z0-9]+`)

const timeFormat = "Mon Jan 2 15:04:05 -0700 2006"

func tweetToNote(t tweet, mediaMap map[string]string, outputDir string, selfUserID string) error {
	if t.Tweet.InReplyTo != "" && t.Tweet.InReplyTo != selfUserID {
		return nil
	}
	if len(t.Tweet.FullText) > 2 && t.Tweet.FullText[0:3] == "RT " {
		return nil
	}
	fm := frontMatter{
		Tags: []string{"imported", "from-twitter"},
	}

	postDir := path.Join(outputDir, t.Tweet.ID)
	if err := os.MkdirAll(postDir, 0750); err != nil {
		return err
	}

	if t.Tweet.InReplyToStatus != "" {
		fm.InReplyTo = path.Join("..", t.Tweet.InReplyToStatus)
	}

	text := escapeMarkdown(t.Tweet.FullText)

	// Retrieve twitpic images
	text = twitpic.ReplaceAllStringFunc(text, func(url string) string {
		fName, err := untwitpic(url, postDir)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to retrieve %s: %v\n", url, err)
			return url
		}
		return fmt.Sprintf(`{{< imgorvid src="%s" >}}`, fName)
	})

	text = urlShortener.ReplaceAllStringFunc(text, func(url string) string {
		realURL, err := unshorten(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to unshorten %s: %v\n", url, err)
			realURL = url
		}
		return fmt.Sprintf("[%s](%s)", realURL, realURL)
	})

	// Add hashtags
	// Adding a newline at the front then removing it to allow for hashtag lookups at the start of tweets
	text = hashtag.ReplaceAllStringFunc("\n"+text, func(s string) string {
		before := s[:1]
		prefix := s[1:2]
		label := s[2:]

		if name, ok := tagMap[label]; ok {
			tag := strings.ToLower(name)
			fm.Tags = append(fm.Tags, tag)
			return fmt.Sprintf("%s{{< friend \"%s\" >}}", before, name)
		} else if prefix == "@" {
			return fmt.Sprintf("%s[@%s](https://twitter.com/%s)", before, label, label)
		} else if prefix == "#" {
			fm.Tags = append(fm.Tags, label)
			return fmt.Sprintf("%s[%s](/tags/%s)", before, label, label)
		} else {
			return s
		}
	})[1:]

	// Add media
	mp := getMediaPaths(t)
	i := 1
	for twitURL, zipPath := range mp {
		ext := path.Ext(zipPath)
		hugoName := fmt.Sprintf("media-%d%s", i, ext)
		hugoMedia := path.Join(postDir, hugoName)

		mediaMap[zipPath] = hugoMedia
		text = strings.Replace(
			text,
			fmt.Sprintf("[%s](%s)", twitURL, twitURL),
			fmt.Sprintf(`{{< imgorvid src="%s" >}}`, hugoName),
			1,
		)
		i++
	}

	// Add date
	publishedDate, err := time.Parse(timeFormat, t.Tweet.CreatedAt)
	if err != nil {
		return err
	}
	fm.Date = publishedDate.Format(time.RFC3339)

	// Add geo ingo if present
	if t.Tweet.Coordinates.Type == "Point" && len(t.Tweet.Coordinates.LonLat) == 2 {
		loc, err := parseLatLon(t.Tweet.Coordinates.LonLat)
		if err == nil {
			fm.Location = loc
		}
	}

	// Create post
	hugoPost, err := os.Create(path.Join(postDir, "index.md"))
	if err != nil {
		return err
	}

	fmt.Fprintln(hugoPost, "---")

	if err := yaml.NewEncoder(hugoPost).Encode(fm); err != nil {
		return err
	}

	fmt.Fprintln(hugoPost, "---")
	fmt.Fprintln(hugoPost, text)

	return nil
}

var markdownEscapable = regexp.MustCompile(`([!\[\]\(\)])`)

func escapeMarkdown(str string) string {
	text := strings.ReplaceAll(str, "\n", "\\\n")
	return markdownEscapable.ReplaceAllString(text, `\$1`)
}

func copyMedia(zf *zip.ReadCloser, mediaMap map[string]string) error {
	for _, f := range zf.File {
		dst, ok := mediaMap[f.Name]
		if !ok {
			continue
		}

		mf, err := f.Open()
		if err != nil {
			return err
		}
		defer doClose(mf, "Unable to close media file within archive")

		mediaFile, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer doClose(mediaFile, "Unable to close media file in blog archive")

		io.Copy(mediaFile, mf)
	}
	return nil
}

func readUntil(r io.Reader, delim string, max int) error {
	success := 0
	for i := 0; i < max; i++ {
		buf := make([]byte, 1)
		if n, err := r.Read(buf); err != nil || n != 1 {
			return err
		}

		if buf[0] == delim[success] {
			success++
			if success == len(delim) {
				return nil
			}
		} else {
			success = 0
		}
	}

	return fmt.Errorf("unable to find delimiter %s in the first %d of the stream", delim, max)
}

var noRedirects = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func unshorten(url string) (string, error) {
	url = strings.Replace(url, "http://", "https://", 1)
	id := path.Base(url)
	host := path.Base(path.Dir(url))
	file := path.Join("shorteners", host, id)
	if err := os.MkdirAll(path.Dir(file), 0755); err != nil {
		return "", err
	}

	if data, err := os.ReadFile(file); err == nil {
		if len(data) == 0 {
			return "", fmt.Errorf("%s is a dead URL", url)
		}
		return string(data), nil
	}

	res, err := noRedirects.Head(url)
	if err != nil {
		return "", err
	}

	loc := res.Header.Get("location")
	if err := os.WriteFile(file, []byte(loc), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to cache %s: %v\n", url, err)
	}
	if loc == "" {
		return "", fmt.Errorf("no Location header")
	}

	return loc, nil
}

func untwitpic(url string, postDir string) (string, error) {
	id := path.Base(url)
	fName := id
	fPath := path.Join(postDir, fName)
	if _, err := os.Stat(fPath); err == nil {
		return id, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("got %d from %s", res.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	imgEl := doc.Find("img").First()
	src, ok := imgEl.Attr("src")
	if !ok {
		return "", fmt.Errorf("no image on twitpic to retrieve URL of")
	}

	resImg, err := http.Get(src)
	if err != nil {
		return "", err
	}
	defer resImg.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("got %d from %s", res.StatusCode, url)
	}

	img, err := os.Create(fPath)
	if err != nil {
		return "", err
	}
	defer doClose(img, "Unable to close image file")
	if _, err := io.Copy(img, resImg.Body); err != nil {
		return "", err
	}

	return fName, nil
}

func getMediaPaths(t tweet) map[string]string {
	paths := make(map[string]string)
	for _, m := range t.Tweet.Entities.Media {
		partName := path.Base(m.MediaURL)
		if strings.Contains(m.MediaURL, "/tweet_video_thumb/") {
			partName = strings.Replace(partName, ".jpg", ".mp4", 1)
		}
		fname := fmt.Sprintf("%s-%s", t.Tweet.ID, partName)

		paths[m.ExpandedURL] = path.Join("data", "tweets_media", fname)
	}
	return paths
}

func parseLatLon(lonLat []string) (fmLoc, error) {
	lat, err := strconv.ParseFloat(lonLat[1], 64)
	if err != nil {
		return fmLoc{}, err
	}
	lon, err := strconv.ParseFloat(lonLat[0], 64)
	if err != nil {
		return fmLoc{}, err
	}

	return fmLoc{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}
