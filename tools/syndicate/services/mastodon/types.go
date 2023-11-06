package mastodon

import (
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/mattn/go-mastodon"
)

type service struct {
	masto    *mastodon.Client
	username string
	password string
}

var _ shared.Service = (*service)(nil)
