package main

import (
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/backfeeder"
	"github.com/by-jp/www.byjp.me/tools/syndicate/poster"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/mmcdole/gofeed"
)

func parseFeed(urlToPath func(string) string, feedReader io.Reader, tagMatcher *regexp.Regexp, syndicationMatchers map[string]*regexp.Regexp) ([]string, poster.ToPostList, backfeeder.ToBackfeedList, error) {
	fp := gofeed.NewParser()
	feed, err := fp.Parse(feedReader)
	if err != nil {
		return nil, nil, nil, err
	}

	toPost := make(poster.ToPostList)
	toBackfeed := make(backfeeder.ToBackfeedList)
	services := make(map[string]struct{})

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
				services[sID.Source] = struct{}{}
				if err != nil {
					return nil, nil, nil, err
				}
				continue
			}

			for sName, bf := range syndicationMatchers {
				if bf.MatchString(ext.Value) {
					toBackfeed[ext.Value] = backfeeder.BackfeedRef{
						Source:   sName,
						LocalURL: item.Link,
					}
					services[sName] = struct{}{}
					break
				}
			}
		}
	}

	var servicesList []string
	for sName := range services {
		servicesList = append(servicesList, sName)
	}

	return servicesList, toPost, toBackfeed, nil
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
