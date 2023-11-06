package main

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/mmcdole/gofeed"
)

type toPostMap map[shared.SyndicationID]shared.Post
type toBackfeedMap map[string]string

func parseFeed(urlToPath func(string) string, feedReader io.Reader, tagMatcher *regexp.Regexp) ([]string, toPostMap, toBackfeedMap, error) {
	fp := gofeed.NewParser()
	feed, err := fp.Parse(feedReader)
	if err != nil {
		return nil, nil, nil, err
	}

	toPost := make(toPostMap)
	toBackfeed := make(toBackfeedMap)
	var services []string

	for _, item := range feed.Items {
		if item.Extensions == nil || item.Extensions["dc"] == nil || item.Extensions["dc"]["relation"] == nil {
			continue
		}

		for _, ext := range item.Extensions["dc"]["relation"] {
			syndicateTag := tagMatcher.FindStringSubmatch(ext.Value)
			if syndicateTag != nil {
				sID := shared.SyndicationID{Source: syndicateTag[1], ID: syndicateTag[2]}
				if _, ok := toPost[sID]; ok {
					continue
				}
				toPost[sID], err = itemToPost(item, urlToPath)
				services = append(services, sID.Source)
				if err != nil {
					return nil, nil, nil, err
				}
				continue
			}

			for _, bf := range []regexp.Regexp{} {
				if bf.MatchString(ext.Value) {
					toBackfeed[ext.Value] = item.Link
					// TODO: Login for backfed posts
					break
				}
			}
		}
	}

	return services, toPost, toBackfeed, nil
}

func itemToPost(item *gofeed.Item, urlToPath func(string) string) (shared.Post, error) {
	p := shared.Post{
		Title:   item.Title,
		URL:     item.Link,
		Summary: item.Description,
		Content: item.Content,
	}

	for _, enc := range item.Enclosures {
		if !strings.HasPrefix(enc.Type, "image/") {
			continue
		}
		img, err := os.ReadFile(urlToPath(enc.URL))
		if err != nil {
			return p, err
		}

		p.Images = append(p.Images, img)
	}

	return p, nil
}
