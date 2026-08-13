package imageservice

import (
	"bytes"
	"image"
	"image/png"

	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"
)

func (s *ImageService) convertToPNG(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {
	var output bytes.Buffer

	if err := png.Encode(&output, img); err != nil {
		return nil, err
	}
	return common.BuildConvertedFile(
		output,
		req,
		"image/png",
	), nil
}
