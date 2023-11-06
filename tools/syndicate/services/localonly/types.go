package localonly

import "github.com/by-jp/www.byjp.me/tools/syndicate/shared"

type service struct {
	dir string
}

var _ shared.Service = (*service)(nil)
