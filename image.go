package streamdeck

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

// Image Generate new base64 png image string from image.Image.
func Image(i image.Image) (string, error) {
	var b bytes.Buffer

	bw := bufio.NewWriter(&b)
	if _, err := bw.WriteString("data:image/png;base64,"); err != nil {
		return "", err
	}

	w := base64.NewEncoder(base64.StdEncoding, bw)
	if err := png.Encode(w, i); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	if err := bw.Flush(); err != nil {
		return "", err
	}

	return b.String(), nil
}

// ImageJpeg Generate new base64 jpeg image string from image.Image.
// quality is 1-100 and recommended to be 85.
func ImageJpeg(i image.Image, quality int) (string, error) {
	var b bytes.Buffer

	bw := bufio.NewWriter(&b)
	if _, err := bw.WriteString("data:image/jpg;base64,"); err != nil {
		return "", err
	}

	w := base64.NewEncoder(base64.StdEncoding, bw)
	if err := jpeg.Encode(w, i, &jpeg.Options{Quality: quality}); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	if err := bw.Flush(); err != nil {
		return "", err
	}

	return b.String(), nil
}

// ImageSvg Generate new svg image string from string.
// Just prepend "data:image/svg+xml;charset=utf8," to the string.
func ImageSvg(s string) string {
	return "data:image/svg+xml;charset=utf8," + s
}
