package localonly

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

func init() {
	shared.Register("localonly", New)
}

func New(cfg map[string]any) (shared.Service, error) {
	dir, ok := cfg["dir"].(string)
	if !ok {
		return nil, fmt.Errorf("missing 'dir' (the directory to store local only posts)")
	}
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to get absolute path for 'dir' %s: %w", dir, err)
	}
	svc := &service{dir: absDir}

	s, err := os.Stat(svc.dir)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("unable to create 'dir' %s: %w", svc.dir, err)
		}

		return svc, nil
	} else if err != nil {
		return nil, fmt.Errorf("invalid 'dir': %w", err)
	} else if !s.IsDir() {
		return nil, fmt.Errorf("invalid 'dir' %s: is not a directory", svc.dir)
	}

	if erase, ok := cfg["erase"].(bool); ok && erase {
		if err := os.RemoveAll(dir); err != nil {
			return nil, fmt.Errorf("unable to erase 'dir' %s: %w", svc.dir, err)
		}

		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("unable to recreate 'dir' %s: %w", svc.dir, err)
		}
	}

	return svc, nil
}

func (s *service) BackfeedMatcher() (*regexp.Regexp, error) {
	return regexp.Compile(`file://` + s.dir + `/(.*)`)
}

func (s *service) Connect(bool) error { return nil }
func (s *service) Close() error       { return nil }
