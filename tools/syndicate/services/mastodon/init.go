package mastodon

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/config"
	"github.com/mattn/go-mastodon"
)

func init() {
	shared.Register("mastodon", New)
	shared.Register("pixelfed", New)
}

func New(cfg map[string]any) (shared.Service, error) {
	cfgMap, err := config.GetStrings(cfg, "server", "username", "access_token_envvar")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(cfgMap["server"])
	if err != nil {
		return nil, err
	}

	srv := &service{
		masto: mastodon.NewClient(&mastodon.Config{
			Server:      u.String(),
			AccessToken: cfgMap["access_token"],
		}),
		username: cfgMap["username"],
	}

	return srv, nil
}

func (s *service) BackfeedMatcher() (*regexp.Regexp, error) {
	url := s.masto.Config.Server
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	return regexp.Compile(
		regexp.QuoteMeta(url) +
			`(?P<username>@[a-zA-Z0-9_-]+)/(?P<id>[0-9]+)/`,
	)
}

func (s *service) Connect(force bool) error {
	if s.masto.Config.AccessToken != "" && !force {
		return nil
	}
	return s.masto.Authenticate(context.Background(), s.username, s.password)
}

func (s *service) Close() error {
	return fmt.Errorf("not implemented")
}
