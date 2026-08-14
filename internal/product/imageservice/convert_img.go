package imageservice

import (
	"errors"

	"go-practice/CONVERTER/internal/converter/models"
)

func (s *ImageService) ConvertIMG(req models.ConvertRequest) (*models.ConvertedFile, error) {
	img, err := s.decodeImage(req.File)
	if err != nil {
		return nil, err
	}

	switch req.OutputFormat {

	case models.FormatPNG:
		return s.convertToPNG(img, req)

	case models.FormatJPG, models.FormatJPEG:
		return s.convertToJPEG(img, req)

	case models.FormatGIF:
		return s.convertToGIF(img, req)

	case models.FormatWEBP:
		return s.convertToWEBP(img, req)

	case models.FormatICO:
		return s.convertToICO(img, req)

	case models.FormatPDF:
		return s.convertToPDF(img, req)

	default:
		return nil, errors.New("unsupported output format")
	}
}
