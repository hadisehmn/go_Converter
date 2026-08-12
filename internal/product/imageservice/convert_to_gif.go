package imageservice

import (
	"bytes"
	"go-practice/CONVERTER/internal/converter/common"
	"go-practice/CONVERTER/internal/converter/models"
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

	return common.BuildConvertedFile(output, req, "image/gif"), nil
}
