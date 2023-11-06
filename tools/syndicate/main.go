package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/poster"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
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
	// bkfd := backfeeder.New(cfg.services)

	for _, feed := range cfg.feeds {
		f, err := os.Open(feed)
		check(err)
		services, toPost, _, err := parseFeed(cfg.urlToPath, f, cfg.tagMatcher)
		check(err)

		fmt.Fprintf(os.Stderr, "Found %d syndications to complete in %s\n", len(toPost), feed)
		fmt.Fprintf(os.Stderr, "Connecting to %s to syndicate…\n", strings.Join(services, ", "))
		for _, sname := range services {
			if err := pstr.Connect(sname); err != nil {
				check(fmt.Errorf("couldn't connect to %s: %w", sname, err))
			}
		}

		for k, p := range toPost {
			if err := pstr.Post(k, p); err == nil {
				fmt.Printf("Posted '%s' to %s: %s\n", p.Title, k.Source, pstr.PostedURL(k))
			} else {
				fmt.Fprintf(os.Stderr, "Couldn't post %s to %s: %v\n", p.URL, k.Source, err)
			}
		}

		for _, fname := range cfg.content {
			f, err := os.ReadFile(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't read %s: %v\n", fname, err)
				continue
			}

			tags := cfg.tagMatcher.FindAllSubmatch(f, -1)
			if tags == nil {
				continue
			}

			var urls []string
			for _, tag := range tags {
				sid := shared.SyndicationID{Source: string(tag[1]), ID: string(tag[2])}
				if url := pstr.PostedURL(sid); url != "" {
					urls = append(urls, url)
					f = bytes.ReplaceAll(f, sid.Bytes(), []byte(url))
					log.Printf("Replacing syndication tag '%s' with post URL: %s", sid, url)
				}
			}

			if len(urls) == 0 {
				continue
			}

			if err := os.WriteFile(fname, f, 0644); err != nil {
				fmt.Fprintf(
					os.Stderr,
					"Couldn't insert posted URLs (%s) into %s: %v\n",
					strings.Join(urls, ", "),
					fname,
					err,
				)
			}
		}
	}
}
