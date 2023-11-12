package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/backfeeder"
	"github.com/by-jp/www.byjp.me/tools/syndicate/poster"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "🛑 %+v\n", err)
		// TODO: Remove panic
		panic(err)
		os.Exit(1)
	}
}

func main() {
	check(os.Chdir("../../"))
	cfg, err := parseConfig(".syndicate")
	check(err)

	pstr := poster.New(cfg.services)
	bkfd := backfeeder.New(cfg.services)

	for _, feed := range cfg.feeds {
		f, err := os.Open(feed)
		check(err)
		defer f.Close()
		services, toPost, toBackfeed, err := parseFeed(cfg.urlToPath, f, cfg.tagMatcher, cfg.syndicationMatchers)
		check(err)

		if len(services) > 0 {
			fmt.Fprintf(os.Stderr, "Connecting to %s to syndicate & backfeed…\n", strings.Join(services, ", "))
			for _, sname := range services {
				if err := pstr.Connect(sname); err != nil {
					check(fmt.Errorf("couldn't connect to %s: %w", sname, err))
				}
			}
		}

		fmt.Fprintf(os.Stderr, "Found %d new syndications to post in %s\n", len(toPost), feed)
		for k, p := range toPost {
			if err := pstr.Post(k, p); err == nil {
				fmt.Printf("Posted '%s' to %s: %s\n", p.Title, k.Source, pstr.PostedURL(k))
			} else {
				fmt.Fprintf(os.Stderr, "Couldn't post %s to %s: %v\n", p.URL, k.Source, err)
			}
		}

		for _, fname := range cfg.content {
			if err := pstr.ReplaceReferences(fname, cfg.tagMatcher); err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't replace syndication references: %v\n", err)
			}
		}

		fmt.Fprintf(os.Stderr, "Found %d existing syndications to backfeed from %s\n", len(toBackfeed), feed)
		_ = bkfd
	}
}
