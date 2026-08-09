package pdfservice

import (
	"fmt"

	"go-practice/CONVERTER/models"
)

func (s *PDFService) ConvertPDF(req models.ConvertRequest) (*models.ConvertedFile, error) {

	switch req.OutputFormat {

	case "png", "jpg", "jpeg":
		return s.convertToImage(req)

	case "txt":
		return s.convertToText(req)

	default:
		return nil, fmt.Errorf(
			"unsupported PDF conversion: %s",
			req.OutputFormat,
		)
	}
}
