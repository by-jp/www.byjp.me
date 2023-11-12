package shared

import (
	"regexp"
	"time"
)

type ServiceCreator func(map[string]any) (Service, error)

type Service interface {
	BackfeedMatcher() (*regexp.Regexp, error)
	Connect(force bool) error
	Post(Post) (url string, err error)
	Interactions(url string) ([]Interaction, error)
	Close() error
}

type Post struct {
	ServiceType string
	Title       string
	Summary     string
	Content     string
	URL         string
	Images      [][]byte
}

type Interaction struct {
	// eg. Repost is 🔁, Facebook is 👍, Instagram is ♥️, Mastodon is ⭐️, Medium is 👏
	Emoji string `json:"emoji,omitempty"`
	// The URL of the original interaction
	URL string `json:"url"`
	// If there's a comment associated with the interaction
	Comment string `json:"comment,omitempty"`
	// Details of the author
	Author    Author    `json:"author"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

type Author struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Avatar    []byte `json:"-"`
	AvatarURL string `json:"-"`
}
