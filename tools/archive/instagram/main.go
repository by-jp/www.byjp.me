package main

import (
	"archive/zip"
	"crypto/md5"
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v2"

	. "github.com/by-jp/www.byjp.me/tools/shared"
)

const titleLength = 48
const minLength = 16

var tagMap = map[string]string{
	"@andyrobert1729":  "Andy",
	"@artphilm":        "Phil",
	"@beckyfuzzymuzzy": "Becky",
	"@bickertonjane":   "Auntie Jane",
	"@bratpack_ldn":    "The Brat Pack",
	"@buckettafloat":   "Erica",
	"@buckettsails":    "Joe",
	"@chrismhs":        "Chris",
	"@dxcompton":       "Dave",
	"@ericabuckett8":   "Erica",
	"@esther_dr":       "Esther",
	"@hazanj99":        "Jenny",
	"@hazanjon":        "Jon",
	"@helenhs16":       "Mum",
	"@joostsposts":     "Joost",
	"@kaphleenmurthy":  "Kathleen",
	"@le_boyd":         "Leanne",
	"@lydiadr":         "Lydia",
	"@Mr_Bingo":        "Mr. Bingo",
	"@ponkalulu":       "Caitlin",
	"@rosalysbryan":    "Rose",
	"@spagbol_terol":   "Paul",
	"@yvetteedrei":     "Yvette",
}

type location struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type post struct {
	Media             []media
	Title             string
	CreationTimestamp int64 `json:"creation_timestamp"`
}

type media struct {
	URI               string
	CreationTimestamp int64 `json:"creation_timestamp"`
	Title             string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path/to/instagram/username_yyyymmdd.zip> <path/to/hugo>\n", os.Args[0])
	}

	archive := os.Args[1]
	hugo := os.Args[2]
	outputDir := path.Join(hugo, "content", "photos")

	zf, err := zip.OpenReader(archive)
	Check(err, "Unable to open instagram archive")
	defer DoClose(zf, "Unable to close zipfile")

	postCount, mediaMap, err := createPosts(zf, outputDir)
	Check(err, "Unable to create hugo posts for your instagram data")
	// TODO: Rewind zip?
	Check(copyMedia(zf, mediaMap), "Unable to copy media to your hugo blog")

	fmt.Printf("Success! %d Instagram posts (with %d images and videos) were added to your hugo blog.\n", postCount, len(mediaMap))
}

var postsFile = regexp.MustCompile(`\Acontent/posts_(\d+)\.json\z`)

func createPosts(zf *zip.ReadCloser, outputDir string) (int, map[string]string, error) {
	for _, f := range zf.File {
		match := postsFile.FindStringSubmatch(f.Name)
		if len(match) == 0 {
			continue
		}

		jf, err := f.Open()
		if err != nil {
			return 0, nil, err
		}
		defer DoClose(jf, "Unable to close posts file within archive")

		return postsFromFile(jf, outputDir)
	}

	return 0, nil, errors.New("no content/posts_1.json file found in zip file")
}

func postsFromFile(r io.Reader, outputDir string) (int, map[string]string, error) {
	postsCount := 0
	mediaMap := make(map[string]string)
	dec := json.NewDecoder(r)

	// Opening [
	tok, err := dec.Token()
	if err != nil {
		return 0, nil, err
	}
	if fmt.Sprintf("%s", tok) != "[" {
		fmt.Println(tok)
		return 0, nil, errors.New("posts JSON doesn't start with '['")
	}

	for dec.More() {
		var p post
		if err := dec.Decode(&p); err != nil {
			return postsCount, mediaMap, errors.New("unable to decode JSON")
		}

		if err := postToPost(p, mediaMap, outputDir); err != nil {
			return postsCount, mediaMap, err
		}
		postsCount++
	}

	return postsCount, mediaMap, nil
}

var hashtag = regexp.MustCompile(`[#@]\w+`)

func postToPost(p post, mediaMap map[string]string, outputDir string) error {
	id, err := postHash(p)
	if err != nil {
		return err
	}
	postDir := path.Join(outputDir, id)

	if err := os.MkdirAll(postDir, 0750); err != nil {
		return err
	}

	text := p.Title
	if text == "" {
		text = p.Media[0].Title
	}
	if text, err = fixEncoding(text); err != nil {
		return err
	}

	fm := FrontMatter{
		Tags: []string{"imported", "from-instagram"},
	}

	fm.Title = text
	if len(fm.Title) > titleLength {
		newEnd := minLength
		for i := titleLength; i > minLength; i-- {
			if fm.Title[i] == ' ' {
				newEnd = i
				break
			}
		}

		fm.Title = fm.Title[:newEnd] + "…"

		if idx := strings.Index(fm.Title, "\n"); idx != -1 {
			fm.Title = fm.Title[:idx]
		}
	}

	text = escapeMarkdown(text)

	// Add hashtags
	text = hashtag.ReplaceAllStringFunc(text, func(s string) string {
		if name, ok := tagMap[s]; ok {
			tag := strings.ToLower(name)
			fm.Tags = append(fm.Tags, tag)
			return fmt.Sprintf("[%s](/tags/%s)", name, tag)
		} else if s[:1] == "@" {
			fmt.Println(s)
			return fmt.Sprintf("[%s](https://instagram.com/%s)", s, s[1:])
		} else {
			fm.Tags = append(fm.Tags, s[1:])
			return fmt.Sprintf("[%s](/tags/%s)", s, s[1:])
		}
	})

	// Add media
	for i, m := range p.Media {
		ext := path.Ext(m.URI)
		if ext == "" {
			// Some video files have their extensions omitted
			ext = ".mp4"
		}

		hugoName := fmt.Sprintf("media-%d%s", i, ext)
		hugoMedia := path.Join(postDir, hugoName)

		fm.Media = append(fm.Media, hugoName)
		mediaMap[m.URI] = hugoMedia
	}

	// Add date
	publishedAt := time.Unix(p.Media[0].CreationTimestamp, 0).UTC()
	fm.Date = publishedAt.Format(time.RFC3339)

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

var fixEncoding = charmap.ISO8859_1.NewEncoder().String

var markdownEscapable = regexp.MustCompile(`([!\[\]\(\)])`)

func escapeMarkdown(str string) string {
	text := strings.ReplaceAll(str, "\n", "\n\n")
	return markdownEscapable.ReplaceAllString(text, `\$1`)
}

func postHash(p post) (string, error) {
	h := md5.New()
	enc := json.NewEncoder(h)
	if err := enc.Encode(p); err != nil {
		return "", err
	}
	b64 := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(h.Sum(nil))
	return strings.ToLower(b64), nil
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
		defer DoClose(mf, "Unable to close media file within archive")

		mediaFile, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer DoClose(mediaFile, "Unable to close media file in blog archive")

		io.Copy(mediaFile, mf)

		if strings.HasSuffix(dst, ".mp4") {
			makeThumbnail(dst, dst+".jpg")
		}
	}
	return nil
}

func makeThumbnail(video, output string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", video, "-vf", "select=eq(n\\,0)", "-q:v", "3", output)
	fmt.Println(cmd.String())
	return cmd.Run()
}
