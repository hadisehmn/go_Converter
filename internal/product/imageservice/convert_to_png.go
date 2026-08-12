package imageservice

import (
	"bytes"
	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"
	"image"

	"github.com/disintegration/imaging"
)

func (s *ImageService) convertToPNG(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	err := imaging.Encode(&output, img, imaging.PNG)
	if err != nil {
		return nil, err
	}

	return common.BuildConvertedFile(output, req, "image/png"), nil
}
