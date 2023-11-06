package instagram

import (
	"fmt"

	"github.com/Davincible/goinsta/v3"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/images"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/text"
)

func (s *service) Post(p shared.Post) (string, error) {
	if len(p.Images) == 0 {
		return "", fmt.Errorf("no images to post to instagram for %s", p.URL)
	}

	jpgIOs, err := images.ToJPEGs(p.Images)
	if err != nil {
		return "", fmt.Errorf("failed to convert images to JPEGs: %w", err)
	}

	upload := &goinsta.UploadOptions{
		Caption: text.Caption(p),
	}
	if len(jpgIOs) == 1 {
		upload.File = jpgIOs[0]
	} else {
		upload.Album = jpgIOs
	}

	item, err := s.insta.Upload(upload)
	if err != nil {
		return "", err
	}

	return postURL(item), nil
}

func postURL(item *goinsta.Item) string {
	return fmt.Sprintf("https://www.instagram.com/p/%s/", item.Code)
}
