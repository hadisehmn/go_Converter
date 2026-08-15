package converter

import (
	"errors"

	"go-practice/CONVERTER/internal/converter/models"
	imageservice "go-practice/CONVERTER/internal/product/imageservice"
	pdfservice "go-practice/CONVERTER/internal/product/pdfservice"
)

type ConvertService struct {
	imageService *imageservice.ImageService
	pdfService   *pdfservice.PDFService
}

func NewConvertService() *ConvertService {
	return &ConvertService{
		imageService: imageservice.NewImageService(),
		pdfService:   pdfservice.NewPDFService(),
	}
}
func isImage(format models.FileFormat) bool {
	switch format {
	case models.FormatPNG,
		models.FormatJPG,
		models.FormatJPEG,
		models.FormatGIF,
		models.FormatICO,
		models.FormatWEBP:

		return true

	default:
		return false
	}
}

func (s *ConvertService) Convert(req models.ConvertRequest) (*models.ConvertedFile, error) {

	switch {
	case isImage(req.InputFormat):
		return s.imageService.ConvertIMG(req)

	case req.InputFormat == models.FormatPDF:
		return s.pdfService.ConvertPDF(req)

	default:
		return nil, errors.New("unsupported input format")
	}
}
