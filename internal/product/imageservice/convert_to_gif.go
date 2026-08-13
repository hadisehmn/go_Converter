package imageservice

import (
	"bytes"
	"image"
	"image/gif"

	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"
)

func (s *ImageService) convertToGIF(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {
	var output bytes.Buffer

	if err := gif.Encode(&output, img, nil); err != nil {
		return nil, err
	}
	return common.BuildConvertedFile(
		output,
		req,
		"image/gif",
	), nil
}
