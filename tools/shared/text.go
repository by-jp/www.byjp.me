package shared

import (
	"bufio"
	"embed"
	"math/big"
	"strings"
)

//go:embed emoji-data.txt
var edf embed.FS
var emojiCodePoints map[rune]struct{}

func init() {
	f, err := edf.Open("emoji-data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	emojiCodePoints = make(map[rune]struct{})

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) < 2 {
			continue
		}

		codeToRune := func(str string) rune {
			cp := new(big.Int)
			cp.SetString(str, 16)
			return rune(cp.Int64())
		}

		codepoints := strings.Split(strings.TrimSpace(parts[0]), "..")
		switch len(codepoints) {
		case 1:
			emojiCodePoints[codeToRune(codepoints[0])] = struct{}{}
		case 2:
			a := codeToRune(codepoints[0])
			b := codeToRune(codepoints[1])
			for i := a; i <= b; i++ {
				emojiCodePoints[i] = struct{}{}
			}
		default:
			panic("Unknown emoji-data.txt database")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// NB. emoji-data.txt needs to be up to date from https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-data.txt
func ExtractLeadingEmoji(str string) (string, string) {
	parts := strings.SplitN(str, " ", 2)
	if len(parts) == 1 {
		return "", str
	}

	for _, c := range parts[0] {
		if _, ok := emojiCodePoints[c]; !ok {
			return "", str
		}
	}

	return parts[0], strings.TrimSpace(parts[1])
}
