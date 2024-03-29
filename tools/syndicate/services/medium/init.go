package medium

import (
	"fmt"
	"regexp"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

func init() {
	shared.Register("medium", New)
}

func New(config map[string]any) (shared.Service, error) {
	return &service{}, nil
}

func (s *service) BackfeedMatcher() (*regexp.Regexp, error) {
	return regexp.Compile(`https://medium.com/(?P<username>[a-zA-Z0-9_-]+)/(?P<id>[a-zA-Z0-9_-]+)/`)
}

func (s *service) Connect(force bool) error {
	return fmt.Errorf("not implemented")
}

func (s *service) Close() error {
	return fmt.Errorf("not implemented")
}
