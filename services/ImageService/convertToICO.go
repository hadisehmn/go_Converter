package imageservice

import (
	"bytes"
	"image"

	"go-practice/CONVERTER/models"

	ico "github.com/Kodeworks/golang-image-ico"
)

func (s *ImageService) convertToICO(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	err := ico.Encode(&output, img)
	if err != nil {
		return nil, err
	}

	return s.buildConvertedFile(output, req, "image/x-icon"), nil
}
