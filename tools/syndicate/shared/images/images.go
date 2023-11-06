package images

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"

	_ "golang.org/x/image/webp"
)

func ToJPEG(imgData []byte) (io.Reader, error) {
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := jpeg.Encode(&b, img, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}
	return &b, nil
}

func ToJPEGs(imgsData [][]byte) ([]io.Reader, error) {
	jpgs := make([]io.Reader, len(imgsData))
	for i, img := range imgsData {
		var err error
		jpgs[i], err = ToJPEG(img)
		if err != nil {
			return nil, err
		}
	}
	return jpgs, nil
}
