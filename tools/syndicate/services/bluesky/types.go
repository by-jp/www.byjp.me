package bluesky

import "github.com/by-jp/www.byjp.me/tools/syndicate/shared"

type Config struct {
	Username string
	Password string
}

type service struct {
	configPath string
}

var _ shared.Service = (*service)(nil)
