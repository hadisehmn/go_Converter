package imageservice

import (
	"bytes"
	"go-practice/CONVERTER/models"
	"image"
	"image/gif"

	"github.com/disintegration/imaging"
)

func (s *ImageService) convertToGIF(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var output bytes.Buffer

	imageCopy := imaging.Clone(img)

	err := gif.Encode(&output, imageCopy, nil)
	if err != nil {
		return nil, err
	}

	return s.buildConvertedFile(output, req, "image/gif"), nil
}
