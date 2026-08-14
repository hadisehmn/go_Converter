package pdfservice

import (
	"fmt"

	"go-practice/CONVERTER/internal/converter/models"
)

func (s *PDFService) ConvertPDF(req models.ConvertRequest) (*models.ConvertedFile, error) {
	switch req.OutputFormat {

	case models.FormatPNG,
		models.FormatJPG,
		models.FormatJPEG:
		return s.convertToImage(req)

	case models.FileFormat("TXT"):
		return s.convertToText(req)

	default:
		return nil, fmt.Errorf(
			"unsupported PDF conversion: %s",
			req.OutputFormat,
		)
	}
}
