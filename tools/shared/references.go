package shared

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
)

var cacheLocation = "./.reference-cache"

func init() {
	if !isDir(cacheLocation) {
		Check(os.Mkdir(cacheLocation, 0755), "couldn't create cache directory")
	}
}

var pathRe = regexp.MustCompile(`[^\w\-. ]+`)

var maxRedirects = 5
var withRedirects = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) >= maxRedirects {
			return fmt.Errorf("more than %d redirects", maxRedirects)
		}
		return nil
	},
	Timeout: 2 * time.Second,
}

func GetReferenceWithCache(url string) (Reference, error) {
	ref := Reference{
		URL:  url,
		Date: time.Now(),
	}

	safePath := path.Join(
		cacheLocation,
		pathRe.ReplaceAllString(url, "/"),
		"reference.yaml",
	)
	failPath := safePath + ".fail"
	if err := os.MkdirAll(path.Dir(safePath), 0755); err != nil {
		return ref, err
	}

	if refFile, err := os.ReadFile(safePath); err == nil {
		return ref, yaml.Unmarshal(refFile, &ref)
	}
	if failFile, err := os.ReadFile(failPath); err == nil {
		return ref, fmt.Errorf(string(failFile))
	}

	if err := getRef(&ref); err != nil {
		os.WriteFile(failPath, []byte(err.Error()), 0644)
		return ref, err
	}

	refData, err := yaml.Marshal(ref)
	if err != nil {
		return ref, err
	}

	if err := os.WriteFile(safePath, refData, 0644); err != nil {
		return ref, err
	}

	return ref, nil
}

func getRef(ref *Reference) error {
	res, err := withRedirects.Get(ref.URL)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got HTTP %d from %s", res.StatusCode, ref.URL)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return fmt.Errorf("is %s not HTML", contentType)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	if date := findAttr(doc, "content", `meta[itemprop="datePublished"]`, `meta[property="article:published_time"]`, `meta[property="video:release_date"]`); date != "" {
		ref.Date, _ = time.Parse(time.RFC3339, date)
	}

	oembedURL := findAttr(doc, "href", `link[rel="alternate"][type="application/json+oembed"]`)
	if oembedURL != "" {
		if err := getRefFromOembed(ref, oembedURL); err != nil {
			fmt.Printf("OEmbed failed: %v\n", err)
		} else {
			return nil
		}
	}

	ref.Type = strings.Split(findAttr(doc, "content", `meta[property="og:type"]`), ".")[0]
	ref.Author = findAttr(doc, "content", `meta[itemprop="name"]`)
	ref.Name = findAttr(doc, "content", `meta[property="og:title"]`, `meta[name="twitter:title"]`)
	if ref.Name == "" {
		ref.Name = findAttr(doc, "", `title`)
	}

	return nil
}

func findAttr(doc *goquery.Document, attr string, selectors ...string) string {
	for _, selector := range selectors {
		var val string
		doc.Find(selector).Each(func(i int, s *goquery.Selection) {
			if s == nil || val != "" {
				return
			}

			if attr == "" {
				val = s.Text()
			} else {
				if attrVal, ok := s.Attr(attr); ok {
					val = attrVal
				}
			}
		})

		if val != "" {
			return val
		}
	}

	return ""
}

type oembed struct {
	Title     string `json:"title"`
	Author    string `json:"author_name"`
	Type      string `json:"type"`
	Thumbnail string `json:"thumbnail_url"`
}

func getRefFromOembed(ref *Reference, url string) error {
	res, err := withRedirects.Get(url)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("oembed http status was %d", res.StatusCode)
	}

	var oe oembed
	if err := json.NewDecoder(res.Body).Decode(&oe); err != nil {
		return err
	}

	ref.Author = oe.Author
	ref.Type = oe.Type
	ref.Name = oe.Title

	return nil
}
