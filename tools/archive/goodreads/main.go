package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func check(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n  %v", msg, err)
		os.Exit(1)
	}
}

type closer interface {
	Close() error
}

func doClose(c closer, msg string) {
	check(c.Close(), msg)
}

type postInfo struct {
	Title        string
	Subtitle     string
	Type         string
	Date         string
	Emoji        string
	Draft        bool
	Tags         []string
	Syndications []string

	path    string
	content string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path/to/goodreads_library_export.csv> <path/to/hugo>\n", os.Args[0])
	}

	csvPath := os.Args[1]
	hugo := os.Args[2]
	outputDir := path.Join(hugo, "content", "posts", "reviews", "books")

	f, err := os.Open(csvPath)
	check(err, "Unable to open csv file")
	defer doClose(f, "Unable to close csv file")

	postCount, err := createPosts(csv.NewReader(f), outputDir)
	check(err, "Unable to create hugo posts for your goodreads data")

	fmt.Printf("Success! %d Goodreads reviews were added to your hugo blog.\n", postCount)
}

var requiredFields = []string{
	"ISBN13",
	"Title",
	"Author",
	"Date Read",
	"Date Added",
	"My Review",
}

func createPosts(c *csv.Reader, outputDir string) (int, error) {
	added := 0
	headers := make(map[string]int)

	for {
		record, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return added, err
		}

		if len(headers) == 0 {
			for i, h := range record {
				headers[h] = i
			}

			for _, f := range requiredFields {
				if _, ok := headers[f]; !ok {
					return added, fmt.Errorf("given CSV doesn't have a %s column", f)
				}
			}
			continue
		}

		m := make(map[string]string)
		for _, f := range requiredFields {
			m[f] = record[headers[f]]
		}

		if m["My Review"] == "" {
			continue
		}

		post, err := postFromRow(m)
		if err != nil {
			return added, err
		}

		postPath := path.Join(outputDir, post.path)

		if err := os.MkdirAll(path.Dir(postPath), 0755); err != nil {
			return added, err
		}

		hugoPost, err := os.Create(postPath)
		if err != nil {
			return added, err
		}

		fmt.Fprintln(hugoPost, "---")
		if err := yaml.NewEncoder(hugoPost).Encode(post); err != nil {
			return added, err
		}
		fmt.Fprintln(hugoPost, "---")
		fmt.Fprintln(hugoPost, post.content)
		hugoPost.Close()
		added += 1
	}

	return added, nil
}

func postFromRow(row map[string]string) (postInfo, error) {
	reviewDate, err := time.Parse("2006/01/02", row["Date Read"])
	if err != nil {
		reviewDate, err = time.Parse("2006/01/02", row["Date Added"])
	}
	if err != nil {
		return postInfo{}, err
	}

	isbn := row["ISBN13"][2 : len(row["ISBN13"])-1]
	content := escapeMarkdown(row["My Review"])

	post := postInfo{
		Type:         "review",
		Draft:        false,
		Emoji:        "📖",
		Title:        row["Title"],
		Subtitle:     "A book by " + row["Author"],
		Tags:         []string{"imported", "from-goodreads"},
		Date:         reviewDate.Format(time.RFC3339),
		Syndications: []string{"https://www.goodreads.com/review/show/"},

		content: fmt.Sprintf("\n{{< book \"%s\" >}}\n\n%s", isbn, content),
		path:    path.Join(kebab(row["Author"]), kebab(row["Title"])+".md"),
	}

	return post, nil
}

var markdownEscapable = regexp.MustCompile(`([!\[\]])`)

func escapeMarkdown(str string) string {
	text := strings.ReplaceAll(str, "<br/>", "\n")
	return markdownEscapable.ReplaceAllString(text, `\$1`)
}

var kebabRE = regexp.MustCompile(`[^a-z0-9-]+`)

func kebab(str string) string {
	return kebabRE.ReplaceAllString(
		strings.ToLower(
			strings.ReplaceAll(str, "'", "")),
		"-",
	)
}
