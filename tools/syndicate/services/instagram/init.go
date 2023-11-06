package instagram

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/Davincible/goinsta/v3"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/config"
)

func init() {
	shared.Register("instagram", New)
}

func New(cfg map[string]any) (shared.Service, error) {
	cfgMap, err := config.GetStrings(cfg, "username", "password_envvar", "totp_envvar")
	if err != nil {
		return nil, err
	}

	svc := &service{
		// TODO: Real dir for .goinsta
		configPath: filepath.Join("/tmp", ".goinsta"),
		username:   cfgMap["username"],
		password:   cfgMap["password"],
		totp:       cfgMap["totp"],
	}

	if svc.username == "" {
		return nil, fmt.Errorf("missing Instagram username")
	}

	return svc, nil
}

func (s *service) BackfeedMatcher() (*regexp.Regexp, error) {
	return regexp.Compile(`https://www.instagram.com/p/(?P<code>[a-zA-Z0-9_-]+)/`)
}

func (s *service) Connect(force bool) error {
	if s.insta != nil && !force {
		return nil
	}

	var err error
	s.insta, err = authenticate(
		s.configPath,
		s.username,
		s.password,
		s.totp,
	)

	return err
}

func (s *service) Close() error {
	return s.insta.Export(s.configPath)
}

func authenticate(config, username, password, totp string) (*goinsta.Instagram, error) {
	_, err := os.Stat(config)
	if err == nil {
		insta, err := goinsta.Import(config)
		if err != nil {
			return nil, err
		}

		return insta, insta.OpenApp()
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	insta := goinsta.New(username, password, totp)
	err = insta.Login()
	if err == nil {
		return insta, nil
	}

	if err != goinsta.Err2FARequired {
		return nil, err
	}

	return insta, insta.TwoFactorInfo.Login2FA()
}
