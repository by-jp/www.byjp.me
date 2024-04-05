package main

import (
	"crypto/md5"
	"encoding/base32"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/by-jp/www.byjp.me/tools/shared"
	"golang.org/x/text/encoding/charmap"
)

var replaceCountRE = regexp.MustCompile(`-\d{3}\.zip`)

func findArchives(archiveFile string) ([]string, error) {
	dir := path.Dir(archiveFile)
	file := replaceCountRE.ReplaceAllString(path.Base(archiveFile), "-*.zip")
	files, err := doublestar.Glob(os.DirFS(dir), file)
	if err != nil {
		return nil, err
	}
	for i, f := range files {
		files[i] = path.Join(dir, f)
	}
	return files, nil
}

func makePostPath(postType string, postDate time.Time, postHash string) string {
	return path.Join("content", postType+"s", "facebook", postDate.Format("2006-01"), postHash)
}

var youtubeRE = regexp.MustCompile(`youtu(?:\.be/|be.com/watch\?v=)(.+)`)

func postize(e PostCheckinPhotoOrVideo, matches []string) (shared.Post, shared.MediaMap, bool, error) {
	archiveRoot := matches[1]

	postHash, err := postHash(e)
	if err != nil {
		return shared.Post{}, nil, false, err
	}
	if shared.ShouldIgnoreHash(postHash) {
		return shared.Post{}, nil, true, err
	}

	postDate := time.Unix(int64(e.Timestamp), 0).UTC()
	post := shared.Post{
		Path: makePostPath("note", postDate, postHash),
		FrontMatter: shared.FrontMatter{
			Title: e.Title,
			Date:  postDate.Format(time.RFC3339),
			Tags: []string{
				"imported",
				"from-facebook",
			},
		},
	}

	// Ignore my posts on others' walls
	if strings.HasPrefix(post.FrontMatter.Title, "JP Hastings-Spital wrote on") {
		return post, nil, true, nil
	}

	for _, attach := range e.Attachments {
		linkURLs := attach.Data.GetSubstrings("external_context", "url")
		for _, link := range linkURLs {
			if link == "" {
				continue
			}

			if strings.HasPrefix(link, "http://") {
				link = "https://" + link[7:]
			}

			if !strings.HasPrefix(link, "https://") {
				link = "https://" + link
			}

			if strings.HasPrefix(link, "https://360.io/") || strings.HasPrefix(link, "https://fb.thisismyjam.com") {
				return shared.Post{}, nil, false, err
			}

			post.Path = makePostPath("notes", postDate, postHash)
			post.FrontMatter.Type = "repost"
			post.FrontMatter.RepostOf = link

			ref, err := shared.GetReferenceWithCache(link)
			if err != nil {
				fmt.Printf("Couldn't get reference for %s because %v\n", link, err)
			}
			// Only bother adding if there's at least a name
			if ref.Name != "" {
				post.FrontMatter.References = append(post.FrontMatter.References, ref)
			}

			if strings.HasSuffix(post.FrontMatter.Title, "shared a link.") || strings.HasSuffix(post.FrontMatter.Title, "added a new photo.") {
				post.FrontMatter.Title = ""
			}
			if post.FrontMatter.Title == "" && ref.Name != "" {
				post.FrontMatter.Title = ref.Name
			}
		}
	}

	mm := make(shared.MediaMap)
	var media []shared.MediaItem
	for _, attach := range e.Attachments {
		mediaPaths := attach.Data.GetSubstrings("media", "uri")
		for _, mediaPath := range mediaPaths {

			// Ignore instagram cross-posts
			if strings.Contains(mediaPath, "Instagram") {
				return shared.Post{}, nil, true, nil
			}

			mediaName := fmt.Sprintf("media-%d%s", len(media), path.Ext(mediaPath))
			post.FrontMatter.Media = append(post.FrontMatter.Media, shared.MediaItem{URL: mediaName})
			post.Path = makePostPath("photo", postDate, postHash)
			mm[path.Join(archiveRoot, mediaPath)] = path.Join(post.Path, mediaName)
		}
	}

	postFile, err := formatPostText(e.Data.GetString("post"), e.Tags)
	if err != nil {
		return shared.Post{}, nil, false, err
	}

	post.PostFile = postFile

	if match := youtubeRE.FindStringSubmatch(post.FrontMatter.BookmarkOf); len(match) > 1 {
		post.PostFile = fmt.Sprintf("%s\n\n{{< youtube \"%s\" >}}", post.PostFile, match[1])
	}

	return post, mm, false, nil
}

func formatPostText(post string, tags []Tag) (string, error) {
	postText, err := fixEncoding(post)
	if err != nil {
		return "", err
	}

	names := make([]string, len(tags))
	if len(tags) > 0 {
		for i, tag := range tags {
			name := tag.Name
			if newName, ok := tagMap[name]; ok {
				name = newName
			}
			names[i] = fmt.Sprintf(`{{< friend "%s" >}}`, name)
		}

		postText = fmt.Sprintf("%s\n\n(with %s)", postText, strings.Join(names, ", "))
	}

	return postText, nil
}

var tagMap = map[string]string{
	"Chris Hastings-Spital": "chris",
	"JP Hastings-Spital":    "jp",
	"Robert Shaw":           "hungoverdrawn",
	"Lenny Martin":          "lenny",
	"Helen HS":              "Mum",
}

func postHash(e PostCheckinPhotoOrVideo) (string, error) {
	h := md5.New()
	unique := fmt.Sprintf("%d||%s", e.Timestamp, e.Data.GetString("post"))
	if _, err := h.Write([]byte(unique)); err != nil {
		return "", err
	}
	b64 := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(h.Sum(nil))
	return strings.ToLower(b64), nil
}

var fixEncoding = charmap.ISO8859_1.NewEncoder().String
