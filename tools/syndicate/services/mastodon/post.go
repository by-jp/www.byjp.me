package mastodon

import (
	"context"
	"fmt"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/images"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared/text"
	"github.com/mattn/go-mastodon"
)

func (s *service) Post(p shared.Post) (string, error) {
	ctx := context.Background()

	jpgIOs, err := images.ToJPEGs(p.Images)
	if err != nil {
		return "", fmt.Errorf("failed to convert images to JPEGs: %w", err)
	}

	var mediaIDs []mastodon.ID
	for _, img := range jpgIOs {
		a, err := s.masto.UploadMediaFromReader(ctx, img)
		if err != nil {
			return "", fmt.Errorf("couldn't upload image: %w", err)
		}
		mediaIDs = append(mediaIDs, a.ID)
	}

	post, err := s.masto.PostStatus(ctx, &mastodon.Toot{
		Status:   text.Caption(p),
		MediaIDs: mediaIDs,
	})
	if err != nil {
		return "", fmt.Errorf("couldn't post status: %w", err)
	}

	return post.URL, nil
}
