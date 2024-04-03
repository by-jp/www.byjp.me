package shared

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (mm MediaMap) CopyMedia(zf *zip.ReadCloser, rootDir string) error {
	for _, f := range zf.File {
		dst, ok := mm[f.Name]
		if !ok {
			continue
		}

		mf, err := f.Open()
		if err != nil {
			return err
		}
		defer DoClose(mf, "Unable to close media file within archive")

		mediaFile, err := os.Create(path.Join(rootDir, dst))
		if err != nil {
			return err
		}
		defer DoClose(mediaFile, "Unable to close media file in blog")

		if _, err := io.Copy(mediaFile, mf); err != nil {
			return fmt.Errorf("unable to copy media file: %w", err)
		}

		if strings.HasSuffix(dst, ".mp4") || strings.HasSuffix(dst, ".flv") {
			if err := makeThumbnail(path.Join(rootDir, dst), path.Join(rootDir, dst)+".jpg"); err != nil {
				fmt.Println(err)
			}
		}

		delete(mm, f.Name)
	}

	if len(mm) > 0 {
		var missing string
		for k := range mm {
			missing = k
			break
		}
		return ErrMissingMedia{Example: missing}
	}

	return nil
}

type ErrMissingMedia struct {
	Example string
}

func (e ErrMissingMedia) Error() string {
	return fmt.Sprintf("some media not available in any archive, eg: %s", e.Example)
}

func (mm MediaMap) Merge(newMediaMap MediaMap) {
	for k, v := range newMediaMap {
		mm[k] = v
	}
}

func makeThumbnail(video, output string) error {
	buf := new(bytes.Buffer)
	cmd := exec.Command("ffmpeg", "-n", "-i", video, "-vframes", "1", "-q:v", "3", output)
	cmd.Stderr = buf

	if err := cmd.Run(); err != nil {
		b, _ := io.ReadAll(buf)
		return fmt.Errorf("couldn't create thumbnail: %s", string(b))
	}
	return nil
}
