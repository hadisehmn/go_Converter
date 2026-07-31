package imageservice

import (
	"bytes"
	"go-practice/CONVERTER/models"
	"image"

	"github.com/disintegration/imaging"
)

func (s *ImageService) convertToJPEG(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	err := imaging.Encode(&output, img, imaging.JPEG)
	if err != nil {
		return nil, err
	}

	return s.buildConvertedFile(output, req, "image/jpeg"), nil
}
