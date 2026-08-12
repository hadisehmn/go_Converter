package imageservice

import (
	"bytes"
	"image"

	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"

	ico "github.com/Kodeworks/golang-image-ico"
)

func (s *ImageService) convertToICO(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	err := ico.Encode(&output, img)
	if err != nil {
		return nil, err
	}

	return common.BuildConvertedFile(output, req, "image/x-icon"), nil
}
