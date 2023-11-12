package backfeeder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/by-jp/www.byjp.me/tools/syndicate/services"
	"github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type BackfeedRef struct {
	Source   string
	LocalURL string
}

type ToBackfeedList map[string]BackfeedRef

type backfeeder struct {
	services  *services.List
	done      map[string]struct{}
	urlToPath func(string) string
}

func New(services *services.List, urlToPath func(string) string) *backfeeder {
	return &backfeeder{
		services:  services,
		done:      make(map[string]struct{}),
		urlToPath: urlToPath,
	}
}

func (b *backfeeder) BackfeedAll(toBackfeed ToBackfeedList) error {
	allIAs := make(map[string][]shared.Interaction)

	for remoteURL, ref := range toBackfeed {
		ias, err := b.services.Service(ref.Source).Interactions(remoteURL)
		if err != nil {
			return err
		}

		path := b.urlToPath(ref.LocalURL)
		allIAs[path] = append(allIAs[path], ias...)
	}

	for postDir, ias := range allIAs {
		if err := writeInteractions(postDir, ias); err != nil {
			return fmt.Errorf("couldn't write interactions into %s: %w", postDir, err)
		}
	}
	return nil
}

func writeInteractions(dir string, ias []shared.Interaction) error {
	if err := os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
		return err
	}
	path := dir + ".json"

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)

	split := map[string]interface{}{
		"interactions": ias,
		"reactions":    make(map[string]int),
	}

	for _, ia := range ias {
		if ia.Emoji != "" {
			split["reactions"].(map[string]int)[ia.Emoji]++
		}
		if ia.Comment != "" {
			split["reactions"].(map[string]int)["💬"]++
		}
	}

	return enc.Encode(split)
}
