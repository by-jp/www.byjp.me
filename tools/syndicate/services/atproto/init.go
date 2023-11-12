package bluesky

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/config"

	"github.com/karalabe/go-bluesky"
)

func init() {
	shared.Register("atproto", New)
}

func New(cfg map[string]any) (shared.Service, error) {
	cfgMap, err := config.GetStrings(cfg, "server", "username", "app_key_envvar")
	if err != nil {
		return nil, err
	}

	return &service{
		server:   cfgMap["server"],
		username: cfgMap["username"],
		appKey:   cfgMap["app_key"],
	}, nil
}

func (s *service) BackfeedMatcher() (*regexp.Regexp, error) {
	return regexp.Compile(`https://bsky.app/profile/(?P<username>[a-zA-Z0-9_-]+)/post/(?P<id>[a-zA-Z0-9_-]+)/`)
}

func (s *service) Connect(force bool) error {
	if s.client != nil && !force {
		return nil
	}

	cl, err := bluesky.Dial(context.Background(), s.server)
	if err != nil {
		return err
	}
	s.client = cl

	err = s.client.Login(context.Background(), s.username, s.appKey)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, bluesky.ErrMasterCredentials):
		return fmt.Errorf("master credentials cannot be used, please use an app key: https://bsky.app/settings/app-passwords")
	case errors.Is(err, bluesky.ErrLoginUnauthorized):
		return fmt.Errorf("incorrect username or app key")
	default:
		return err
	}
}

func (s *service) Close() error {
	return s.client.Close()
}
