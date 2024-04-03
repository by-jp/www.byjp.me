package shared

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

var ignoreHashes = map[string]struct{}{}

func init() {
	if ig, err := os.ReadFile(".archiveignore"); err == nil {
		for _, hash := range strings.Split(string(ig), "\n") {
			ignoreHashes[hash] = struct{}{}
		}
	}
}

func ShouldIgnoreHash(hash string) bool {
	_, ok := ignoreHashes[hash]
	return ok
}

func ExtractPostsFromArchives[ArchiveEntry interface{}](
	archives []string,
	postsFileRE *regexp.Regexp,
	hugoDir string,
	postize func(ArchiveEntry, []string) (Post, MediaMap, bool, error),
) error {
	mediaMap := make(MediaMap)

	for i, archive := range archives {
		zf, err := zip.OpenReader(archive)
		if err != nil {
			return fmt.Errorf("unable to open archive (%s): %w", archive, err)
		}
		defer DoClose(zf, "Unable to close archive zipfile")

		readErr := readEntriesFromArchive(zf, postsFileRE, func(ae ArchiveEntry, matches []string) error {
			post, mm, ignore, err := postize(ae, matches)
			if err != nil {
				return err
			}
			if ignore {
				return nil
			}

			if err := post.WriteTo(hugoDir); err != nil {
				return err
			}
			mediaMap.Merge(mm)

			return nil
		})
		if readErr != nil {
			return readErr
		}

		if err := mediaMap.CopyMedia(zf, hugoDir); err != nil {
			_, isMissingMedia := err.(ErrMissingMedia)
			if !isMissingMedia || i >= len(archives)-1 {
				return err
			}
		}
	}
	return nil
}

var ErrNoPostsFile = errors.New("no posts file found in the archive zipfile")

func readEntriesFromArchive[ArchiveEntry interface{}](
	zf *zip.ReadCloser,
	postsFileRE *regexp.Regexp,
	processArchiveEntry func(ArchiveEntry, []string) error,
) error {
	for _, f := range zf.File {
		matches := postsFileRE.FindStringSubmatch(f.Name)
		if len(matches) == 0 {
			continue
		}

		jf, err := f.Open()
		if err != nil {
			return err
		}
		defer DoClose(jf, "Unable to close posts file within archive")

		var entry ArchiveEntry
		entries, err := readEntriesFromArchiveJSON(jf, entry)
		if err != nil {
			return err
		}

		for _, entry := range entries {
			if err := processArchiveEntry(entry, matches); err != nil {
				return err
			}
		}
	}

	return nil
}

func readEntriesFromArchiveJSON[ArchiveEntry interface{}](r io.Reader, entry ArchiveEntry) ([]ArchiveEntry, error) {
	dec := json.NewDecoder(r)

	// Opening [
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if fmt.Sprintf("%s", tok) != "[" {
		return nil, fmt.Errorf("posts JSON starts with '%s', instead of '['", tok)
	}

	var entries []ArchiveEntry
	for dec.More() {
		var e ArchiveEntry
		if err := dec.Decode(&e); err != nil {
			return nil, fmt.Errorf("unable to decode JSON: %w", err)
		}
		entries = append(entries, e)
	}

	return entries, nil
}

func (p Post) WriteTo(hugoDir string) error {
	postRoot := path.Join(hugoDir, p.Path)
	if err := os.MkdirAll(postRoot, 0755); err != nil {
		return err
	}

	hugoPost, err := os.Create(path.Join(postRoot, "index.md"))
	if err != nil {
		return err
	}

	fmt.Fprintln(hugoPost, "---")

	if err := yaml.NewEncoder(hugoPost).Encode(p.FrontMatter); err != nil {
		return err
	}

	fmt.Fprintln(hugoPost, "---")
	fmt.Fprintln(hugoPost, p.PostFile)

	return nil
}
