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
	bkfd := backfeeder.New(cfg.services, cfg.urlToInteractionsPath)

	for _, feed := range cfg.feeds {
		f, err := os.Open(feed)
		check(err)
		defer f.Close()
		services, toPost, toBackfeed, err := parseFeed(cfg.urlToPublicPath, f, cfg.tagMatcher, cfg.syndicationMatchers)
		check(err)

		if len(services) > 0 {
			fmt.Fprintf(os.Stderr, "Connecting to %s to syndicate…\n", strings.Join(services, ", "))
			for _, sname := range services {
				if err := cfg.services.Init(sname); err != nil {
					check(fmt.Errorf("couldn't connect to %s: %w", sname, err))
				}
			}
		}

		fmt.Fprintf(os.Stderr, "Found %d new syndications to post in %s\n", len(toPost), feed)
		if _, err := pstr.PostAll(toPost); err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't post syndications: %v\n", err)
		}

		for _, fname := range cfg.content {
			if err := pstr.ReplaceReferences(fname, cfg.tagMatcher); err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't replace syndication references: %v\n", err)
			}
		}

		if !cfg.performBackfeed {
			continue
		}

		fmt.Fprintf(os.Stderr, "Found %d existing syndications to backfeed from %s\n", len(toBackfeed), feed)
		if err := bkfd.BackfeedAll(toBackfeed); err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't backfeed syndications: %v\n", err)
		}
	}
}
