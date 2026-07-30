package service

import (
	"errors"
	"go-practice/CONVERTER/models"
)

type ConvertService struct {
	imageService    *ImageService
	pdfService      *PDFService
	documentService *DocumentService
}

func NewConvertService() *ConvertService {
	return &ConvertService{
		imageService:    NewImageService(),
		pdfService:      NewPDFService(),
		documentService: NewDocumentService(),
	}
}

func isImage(input string, output string) bool {

	switch input {
	case "jpg", "jpeg", "png":
		switch output {
		case "jpg", "jpeg", "png":
			return true
		}
	}

	return false
}

func (s *ConvertService) Convert(req models.ConvertRequest) (*models.ConvertedFile, error) {

	switch {

	case isImage(req.InputFormat, req.OutputFormat):
		return s.imageService.ConvertIMG(req.File, req.Header, req.OutputFormat)
	default:
		return nil, errors.New("unsupported conversion")
	}
}
