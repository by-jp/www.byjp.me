package instagram

import (
	"fmt"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

func (s *service) Interactions(url string) ([]shared.Interaction, error) {
	u, err := s.insta.Profiles.ByName(s.username)
	if err != nil {
		return nil, fmt.Errorf("unable to load your feed to scan for interactions: %w", err)
	}

	feed := u.Feed()
	fmt.Println(feed, feed.Items, feed.NumResults, feed.Status)

	if err := feed.GetCommentInfo(); err != nil {
		return nil, fmt.Errorf("unable load feed comment info: %w", err)
	}

	return nil, fmt.Errorf("not implemented")
}
