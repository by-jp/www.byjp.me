package bluesky

import (
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/karalabe/go-bluesky"
)

type service struct {
	server   string
	username string
	appKey   string
	client   *bluesky.Client
}

var _ shared.Service = (*service)(nil)
