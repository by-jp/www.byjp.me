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

const someLongText = "Here's some longer text that will stand in for a body of a post. The first paragraph is longer than 140 characters, but it also has additional paragraphs.\n\nLike this one."

func TestExtractSummary(t *testing.T) {
	originalSummary := "The original summary, as provided by the article."
	shortSummary := "Some short summary."

	cases := [][]string{
		{"Short summary", originalSummary, shortSummary + "\n\n" + someLongText, shortSummary, someLongText},
		{"No summary", originalSummary, someLongText, originalSummary, someLongText},
		{"Only summary", originalSummary, shortSummary, shortSummary, shortSummary},
		{"Long original summary, no summary in my text", "original" + someLongText, "mine" + someLongText, "", "mine" + someLongText},
	}

	for _, c := range cases {
		summary, body := shared.ExtractSummary(c[1], c[2])
		assert.Equal(t, c[3], summary)
		assert.Equal(t, c[4], body)
	}
}
