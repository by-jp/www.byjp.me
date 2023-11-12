package mastodon

import (
	"context"
	"fmt"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	"github.com/mattn/go-mastodon"
)

func (s *service) Interactions(url string) ([]shared.Interaction, error) {
	id, err := s.postID(url)
	if err != nil {
		return nil, err
	}

	var ias []shared.Interaction

	ctx := context.Background()
	favs, err := s.masto.GetFavouritedBy(ctx, id, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to get favourites for %s: %w", url, err)
	}

	for _, acc := range favs {
		ias = append(ias, shared.Interaction{
			Emoji: "⭐️",
			Author: shared.Author{
				Name:      acc.DisplayName,
				URL:       acc.URL,
				AvatarURL: acc.Avatar,
			},
			Timestamp: acc.CreatedAt,
		})
	}

	reblogs, err := s.masto.GetRebloggedBy(ctx, id, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to get boosts for %s: %w", url, err)
	}

	for _, acc := range reblogs {
		ias = append(ias, shared.Interaction{
			Emoji: "🔁",
			Author: shared.Author{
				Name:      acc.DisplayName,
				URL:       acc.URL,
				AvatarURL: acc.Avatar,
			},
			Timestamp: acc.CreatedAt,
		})
	}

	postCtx, err := s.masto.GetStatusContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("unable to get replies for %s: %w", url, err)
	}

	for _, reply := range postCtx.Descendants {
		if reply.Visibility != mastodon.VisibilityPublic {
			continue
		}

		ias = append(ias, shared.Interaction{
			URL:     reply.URL,
			Comment: reply.Content,
			Author: shared.Author{
				Name:      reply.Account.DisplayName,
				URL:       reply.Account.URL,
				AvatarURL: reply.Account.Avatar,
			},
			Timestamp: reply.CreatedAt,
		})
	}

	return ias, nil
}

func (s *service) postID(url string) (mastodon.ID, error) {
	re, err := s.BackfeedMatcher()
	if err != nil {
		return "", fmt.Errorf("couldn't create the backfeed matcher: %w", err)
	}

	m := re.FindStringSubmatch(url)
	if len(m) == 0 {
		return "", fmt.Errorf("couldn't extract the post ID from the URL")
	}

	var id string
	for i, name := range re.SubexpNames() {
		if name == "id" {
			id = m[i]
			break
		}
	}

	if id == "" {
		return "", fmt.Errorf("couldn't extract the post ID from the URL")
	}

	return mastodon.ID(id), nil
}
