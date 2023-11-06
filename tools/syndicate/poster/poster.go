package poster

import (
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type poster struct {
	services map[string]shared.Service
	done     map[shared.SyndicationID]string
}

func New(services map[string]shared.Service) *poster {
	return &poster{
		services: services,
		done:     make(map[shared.SyndicationID]string),
	}
}

func (p *poster) Connect(sname string) error {
	return p.services[sname].Connect(false)
}

func (p *poster) Post(sid shared.SyndicationID, post shared.Post) error {
	if _, ok := p.done[sid]; ok {
		return nil
	}

	url, err := p.services[sid.Source].Post(post)
	if err != nil {
		return err
	}

	p.done[sid] = url

	return nil
}

func (p *poster) PostedURL(sid shared.SyndicationID) string {
	if url, ok := p.done[sid]; ok {
		return url
	} else {
		return ""
	}
}
