package imageservice

import (
	"bytes"
	"image"
	"image/jpeg"

	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"
)

func (s *ImageService) convertToJPEG(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {
	var output bytes.Buffer

	if err := jpeg.Encode(&output, img, &jpeg.Options{
		Quality: 90,
	}); err != nil {
		return nil, err
	}
	return common.BuildConvertedFile(
		output,
		req,
		"image/jpeg",
	), nil
}
