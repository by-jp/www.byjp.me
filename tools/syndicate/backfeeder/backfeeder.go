package backfeeder

import (
	"github.com/by-jp/www.byjp.me/tools/syndicate/services"
)

type backfeeder struct {
	services *services.List
	done     map[string]struct{}
}

func New(services *services.List) *backfeeder {
	return &backfeeder{
		services: services,
		done:     make(map[string]struct{}),
	}
}
