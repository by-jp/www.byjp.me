package localonly

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/text"

	"github.com/h2non/filetype"
)

var kebab = regexp.MustCompile(`[^a-z0-9]+`)

func (s *service) Post(p shared.Post) (string, error) {
	id := kebab.ReplaceAllString(strings.ToLower(p.Title), "-")
	dir := filepath.Join(s.dir, id)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("unable to save post: %w", err)
	}

	var data string
	for i, img := range p.Images {
		ft, err := filetype.Match(img)
		if err != nil {
			return "", fmt.Errorf("unable to determine image type: %w", err)
		}
		ext := ft.Extension
		imgName := fmt.Sprintf("%03d.%s", i+1, ext)
		if err := os.WriteFile(filepath.Join(dir, imgName), img, 0644); err != nil {
			return "", fmt.Errorf("unable to save post image %d: %w", i, err)
		}

		if len(data) > 0 {
			data = data + "\n"
		}
		data = data + "![](" + imgName + ")"
	}

	if len(data) > 0 {
		data = data + "\n\n"
	}
	data = data + text.Caption(p)
	if err := os.WriteFile(filepath.Join(dir, "index.md"), []byte(data), 0644); err != nil {
		return "", fmt.Errorf("unable to save post: %w", err)
	}

	return "file://" + dir + "/", nil
}
