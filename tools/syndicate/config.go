package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/by-jp/www.byjp.me/tools/syndicate/services"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"

	"github.com/bmatcuk/doublestar/v4"

	_ "github.com/by-jp/www.byjp.me/tools/syndicate/services/bluesky"
	_ "github.com/by-jp/www.byjp.me/tools/syndicate/services/instagram"
	_ "github.com/by-jp/www.byjp.me/tools/syndicate/services/localonly"
	_ "github.com/by-jp/www.byjp.me/tools/syndicate/services/mastodon"
	_ "github.com/by-jp/www.byjp.me/tools/syndicate/services/medium"
)

type fileConfig struct {
	Feeds           []string       `toml:"feeds"`
	ContentGlob     string         `toml:"content_glob"`
	InteractionsDir string         `toml:"interactions_dir"`
	Sites           map[string]any `toml:"sites"`
	PublishRoot     string         `toml:"publish_root"`
	PublishURL      string         `toml:"publish_url"`
}

type config struct {
	feeds               []string
	services            *services.List
	interactionsDir     string
	content             []string
	tagMatcher          *regexp.Regexp
	syndicationMatchers map[string]*regexp.Regexp
	urlToPath           func(string) string
}

func parseConfig(cfgPath string) (*config, error) {
	var cfgData fileConfig
	_, err := toml.DecodeFile(cfgPath, &cfgData)
	if err != nil {
		return nil, err
	}

	cfg := &config{
		feeds:               []string{},
		syndicationMatchers: make(map[string]*regexp.Regexp),
		interactionsDir:     cfgData.InteractionsDir,
		urlToPath: func(url string) string {
			return path.Join(cfgData.PublishRoot, strings.TrimPrefix(url, cfgData.PublishURL))
		},
	}

	cfg.content, err = doublestar.Glob(os.DirFS("."), cfgData.ContentGlob)
	if err != nil {
		return nil, err
	}

	for _, feed := range cfgData.Feeds {
		newFeeds, err := doublestar.Glob(os.DirFS(cfgData.PublishRoot), feed)
		if err != nil {
			return nil, err
		}
		for _, nf := range newFeeds {
			cfg.feeds = append(cfg.feeds, path.Join(cfgData.PublishRoot, nf))
		}
	}

	var serviceTags []string
	for name, siteConfig := range cfgData.Sites {
		svc, err := cfg.services.Load(name, siteConfig)
		if err != nil {
			return nil, err
		}

		bf, err := svc.BackfeedMatcher()
		if err != nil {
			return nil, fmt.Errorf("cannot perform backfeed matching for '%s' because %w", name, err)
		}
		cfg.syndicationMatchers[name] = bf
		serviceTags = append(serviceTags, name)
	}
	cfg.tagMatcher, err = shared.TagMatcher(serviceTags)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
