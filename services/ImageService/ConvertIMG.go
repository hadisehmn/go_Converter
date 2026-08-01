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

	case "png":
		return s.convertToPNG(img, req)

	case "jpg", "jpeg":
		return s.convertToJPEG(img, req)

	case "gif":
		return s.convertToGIF(img, req)

	case "webp":
		return s.convertToWEBP(img, req)

	case "ico":
		return s.convertToICO(img, req)

	case "pdf":
		return s.convertToPDF(img, req)

	default:
		return nil, errors.New("unsupported output format")
	}
}
