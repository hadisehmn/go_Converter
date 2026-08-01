package imageservice

import (
	"bytes"
	"image"

	"go-practice/CONVERTER/models"

	"github.com/chai2010/webp"
)

func (s *ImageService) convertToWEBP(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	err := webp.Encode(&output, img, &webp.Options{
		Lossless: true,
		Quality:  100,
	})
	if err != nil {
		return nil, err
	}

	return s.buildConvertedFile(output, req, "image/webp"), nil
}
