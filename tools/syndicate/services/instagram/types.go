package instagram

import (
	"github.com/Davincible/goinsta/v3"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type service struct {
	insta      *goinsta.Instagram
	configPath string
	username   string
	password   string
	totp       string
}

var _ shared.Service = (*service)(nil)
