package shared_test

import (
	"testing"

	"github.com/by-jp/www.byjp.me/tools/shared"
	"github.com/stretchr/testify/assert"
)

func TestExtractLeadingEmoji(t *testing.T) {
	cases := [][]string{
		{"No Emoji here", "", "No Emoji here"},
		{"A single character start", "", "A single character start"},
		{"😊 The first emoji", "😊", "The first emoji"},
		{"😊 \t  Extra space", "😊", "Extra space"},
		{"🍋‍🟩 15.1 emoji", "🍋‍🟩", "15.1 emoji"},
		{"絵文字", "", "絵文字"},
	}

	for _, c := range cases {
		emoji, text := shared.ExtractLeadingEmoji(c[0])
		assert.Equal(t, c[1], emoji)
		assert.Equal(t, c[2], text)
	}
}
