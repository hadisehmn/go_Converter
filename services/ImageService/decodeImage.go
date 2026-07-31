package imageservice

import (
	"bytes"
	"errors"
	"image"
	"io"
	"mime/multipart"

	"github.com/disintegration/imaging"
)

func (s *ImageService) decodeImage(file multipart.File) (image.Image, error) {

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	img, err := imaging.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("invalid image")
	}

	return img, nil
}
