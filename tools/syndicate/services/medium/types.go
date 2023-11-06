package medium

import "github.com/by-jp/www.byjp.me/tools/syndicate/shared"

type Config struct {
	Username string
	Password string
}

type service struct {
}

var _ shared.Service = (*service)(nil)
