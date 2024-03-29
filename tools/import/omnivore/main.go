package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"slices"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

//go:embed query.gql
var gql string

var ignoreLabels = []string{
	"opinion-agree",
	"opinion-disagree",
	"interesting",
}

func main() {
	apiKey, ok := os.LookupEnv("OMNIVORE_API_KEY")
	if !ok || len(apiKey) == 0 {
		fmt.Fprint(os.Stderr, "OMNIVORE_API_KEY is not set")
		os.Exit(1)
	}

	rootDir := "./"
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	}

	outputDir := path.Join(rootDir, "content/bookmarks")
	if !isDir(outputDir) {
		fmt.Printf(
			"Usage: %s [directory]\n  'directory' should be the root of your hugo blog (default: ./)\n",
			path.Base(os.Args[0]),
		)
		os.Exit(1)
	}

	// Make the GraphQL request
	articles, err := omnivoreArticles(
		"in:archive has:highlights sort:updated-des",
		apiKey,
	)
	if err != nil {
		fmt.Println("Failed retrieve articles:", err)
		os.Exit(1)
	}

	for _, article := range articles {
		if err := outputArticle(article, outputDir); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to output article: %v\n", err)
		}
	}
}

func isDir(pathStr string) bool {
	st, err := os.Stat(pathStr)
	if os.IsNotExist(err) {
		return false
	}

	return st.IsDir()
}

var hashtags = regexp.MustCompile(`#\w+`)

func outputArticle(article Article, outputDir string) error {
	slug := kebab(article.Title)
	articlePath := path.Join(outputDir, fmt.Sprintf("%s.md", slug))

	fm, _ := loadFrontmatter(articlePath)

	hugoPost, err := os.Create(articlePath)
	if err != nil {
		return err
	}

	if fm.Date == "" {
		fm.Date = article.BookmarkDate.Format(time.RFC3339)
	}

	fm.Title = article.Title
	fm.BookmarkOf = article.OriginalURL
	fm.Tags = removeDupes(append(fm.Tags, article.Tags...))

	if fm.References == nil {
		fm.References = make(map[string]Ref)
	}
	ref := fm.References["bookmark"]
	ref.URL = article.OriginalURL
	ref.Type = "entry"
	ref.Name = article.OriginalTitle
	ref.Summary = article.OriginalSummary
	ref.Author = article.OriginalAuthor
	fm.References["bookmark"] = ref

	if !article.PublishDate.IsZero() {
		fm.PublishDate = article.PublishDate.Format(time.RFC3339)
	}

	fmt.Fprintln(hugoPost, "---")

	if err := yaml.NewEncoder(hugoPost).Encode(fm); err != nil {
		return err
	}

	fmt.Fprint(hugoPost, "---\n\n")
	fmt.Fprintln(hugoPost, linkHashtags(article.Annonation, fm.Tags))

	if len(article.Highlights) > 0 {
		fmt.Fprint(hugoPost, "\n### Highlights\n")
	}

	for i, highlight := range article.Highlights {
		quote := "> " + strings.ReplaceAll(trimQuote(highlight.Quote), "\n", "\n> ")
		fmt.Fprint(hugoPost, "\n"+quote+"\n")

		if highlight.Comment != "" {
			fmt.Fprint(hugoPost, "\n"+linkHashtags(highlight.Comment, fm.Tags)+"\n")
		}

		if i < len(article.Highlights)-1 {
			fmt.Fprint(hugoPost, "\n---\n")
		}
	}

	return nil
}

var allBold = regexp.MustCompile(`\*\*([^*]+)\*\*(\W)?`)

func removeDupes(tags []string) []string {
	slices.Sort(tags)
	return slices.Compact(tags)
}

func trimQuote(quote string) string {
	noTrail := strings.TrimRight(quote, "\n ")
	return allBold.ReplaceAllString(noTrail, "$1$2")
}

func loadFrontmatter(path string) (FrontMatter, error) {
	f, err := os.Open(path)
	if err != nil {
		return FrontMatter{}, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	var fm FrontMatter
	if err := decoder.Decode(&fm); err != nil {
		return FrontMatter{}, err
	}

	return fm, nil
}

func linkHashtags(text string, tags []string) string {
	return hashtags.ReplaceAllStringFunc(text, func(hashtag string) string {
		tags = append(tags, hashtag[1:])
		return fmt.Sprintf("[%s](/tags/%s)", hashtag[1:], strings.ToLower(hashtag[1:]))
	})
}

var kebaber = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func kebab(str string) string {
	return kebaber.ReplaceAllString(strings.ToLower(str), "-")
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

const omnivoreEndpoint = "https://api-prod.omnivore.app/api/graphql"

type Article struct {
	ID              string
	Title           string
	Emoji           string
	BookmarkDate    time.Time
	PublishDate     time.Time
	OriginalTitle   string
	OriginalURL     string
	OriginalSummary string
	OriginalAuthor  string
	Annonation      string
	Highlights      []ArticleHighlight
	Tags            []string
}

type ArticleHighlight struct {
	Quote   string
	Comment string
}

type FrontMatter struct {
	Title       string
	Date        string
	Emoji       string `yaml:"emoji,omitempty"`
	PublishDate string `yaml:"publishDate,omitempty"`
	BookmarkOf  string `yaml:"bookmarkOf"`
	References  map[string]Ref
	Tags        []string
}

type Ref struct {
	URL     string `yaml:"url"`
	Rel     string `yaml:"rel,omitempty"`
	Type    string `yaml:"type,omitempty"`
	Name    string `yaml:"name"`
	Summary string `yaml:"summary,omitempty"`
	Author  string `yaml:"author,omitempty"`
}

type SearchResults struct {
	Data struct {
		Search struct {
			Edges []struct {
				Node SearchResult
			}
			PageInfo struct {
				HasNextPage bool   `json:"hasNextPage"`
				EndCursor   string `json:"endCursor"`
			}
		}
	}
}

type SearchResult struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	OriginalArticleURL string `json:"originalArticleUrl"`
	Author             string `json:"author"`
	PublishedAt        string `json:"publishedAt"`
	ReadAt             string `json:"readAt"`
	Description        string `json:"description"`
	Highlights         []Highlight
	Labels             []struct {
		Name string `json:"name"`
	} `json:"labels"`
}

type Highlight struct {
	Type       string  `json:"type"`
	Position   float64 `json:"highlightPositionPercent"`
	Annotation string  `json:"annotation"`
	Quote      string  `json:"quote"`
}

func omnivoreArticles(query string, apiKey string) ([]Article, error) {
	cursor := ""
	var articles []Article
	for {
		newArticles, nextCursor, err := omnivoreRequest(query, apiKey, cursor)
		if err != nil {
			return nil, err
		}

		articles = append(articles, newArticles...)
		if len(nextCursor) == 0 {
			break
		}
		cursor = nextCursor
	}
	return articles, nil
}

func omnivoreRequest(query, apiKey, cursor string) ([]Article, string, error) {
	request := GraphQLRequest{
		Query: gql,
		Variables: map[string]interface{}{
			"query": query,
			"after": cursor,
		},
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		return nil, "", err
	}

	req, err := http.NewRequest("POST", omnivoreEndpoint, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, "", err
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, "", err
	}

	return parseResponse(body)
}

var titleSplitter = regexp.MustCompile(`\s[-–—|]\s`)

func parseResponse(body []byte) ([]Article, string, error) {
	var searchResults SearchResults
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return nil, "", err
	}

	var articles []Article
	for _, edge := range searchResults.Data.Search.Edges {
		sr := edge.Node

		articleURL := stripMarketing(sr.OriginalArticleURL)

		var highlights []ArticleHighlight
		var annotation string
		for _, highlight := range sr.Highlights {
			if highlight.Type == "NOTE" {
				annotation = highlight.Annotation
			} else {
				highlights = append(highlights, ArticleHighlight{
					Quote:   highlight.Quote,
					Comment: highlight.Annotation,
				})
			}
		}

		if len(annotation) == 0 {
			continue
		}

		bookmarked, err := time.Parse(time.RFC3339, sr.ReadAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse ReadAt date: %s\n", sr.ID)
			continue
		}
		published, err := time.Parse(time.RFC3339, sr.PublishedAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse PublishedAt date (for %s): %s\n", articleURL, sr.ID)
		}

		abbreviatedOriginalTitle := sr.Title
		if parts := titleSplitter.Split(abbreviatedOriginalTitle, -1); len(parts) > 1 {
			abbreviatedOriginalTitle = parts[0]
			for _, part := range parts[1:] {
				if len(part) > len(abbreviatedOriginalTitle) {
					abbreviatedOriginalTitle = part
				}
			}
		}
		title := abbreviatedOriginalTitle
		if annotation[0:2] == "# " {
			parts := strings.SplitN(annotation, "\n", 2)
			title = parts[0][2:]
			annotation = parts[1]
		}

		article := Article{
			ID:              sr.ID,
			Title:           title,
			OriginalTitle:   abbreviatedOriginalTitle,
			OriginalURL:     articleURL,
			OriginalAuthor:  sr.Author,
			OriginalSummary: sr.Description,
			BookmarkDate:    bookmarked,
			PublishDate:     published,
			Highlights:      highlights,
			Annonation:      annotation,
		}

		for _, label := range sr.Labels {
			if slices.Contains(ignoreLabels, label.Name) {
				continue
			}
			article.Tags = append(article.Tags, label.Name)
		}

		articles = append(articles, article)
	}

	var cursor string
	if searchResults.Data.Search.PageInfo.HasNextPage {
		cursor = searchResults.Data.Search.PageInfo.EndCursor
	}

	return articles, cursor, nil
}

func stripMarketing(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse URL: %s\n", rawURL)
		return rawURL
	}

	q := u.Query()
	q.Del("amp")
	q.Del("utm_source")
	q.Del("utm_medium")
	q.Del("utm_campaign")
	q.Del("utm_content")
	q.Del("utm_term")
	u.RawQuery = q.Encode()

	return u.String()
}
