package text

import (
	"strings"

	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
	strip "github.com/grokify/html-strip-tags-go"
)

func StripHTML(s string) string {
	return strings.TrimSpace(strip.StripTags(s))
}

func Caption(p shared.Post) string {
	return p.Title + " · " + StripHTML(p.Summary) + "\n·\nOriginal: " + p.URL
}
