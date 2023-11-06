package backfeeder

import "github.com/by-jp/www.byjp.me/tools/syndicate/shared"

type backfeeder struct {
	services map[string]shared.Service
	done     map[string]struct{}
}

func New(services map[string]shared.Service) *backfeeder {
	return &backfeeder{
		services: services,
		done:     make(map[string]struct{}),
	}
}
