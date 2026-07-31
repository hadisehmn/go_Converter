package imageservice

import (
	"errors"
	"go-practice/CONVERTER/models"
	"strings"
)

func (s *ImageService) ConvertIMG(req models.ConvertRequest) (*models.ConvertedFile, error) {

	img, err := s.decodeImage(req.File)
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(req.OutputFormat) {

	case "jpg", "jpeg":
		return s.convertToJPEG(img, req)

	case "png":
		return s.convertToPNG(img, req)

	default:
		return nil, errors.New("unsupported output format")
	}
}
