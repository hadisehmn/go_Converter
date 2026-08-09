package pdfservice

import (
	"go-practice/CONVERTER/models"
	"io"

	"github.com/ledongthuc/pdf"
)

func (s *PDFService) convertToText(req models.ConvertRequest) (*models.ConvertedFile, error) {

	reader, err := pdf.NewReader(req.File, -1)
	if err != nil {
		return nil, err
	}
	text, err := reader.GetPlainText()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(text)
	if err != nil {
		return nil, err
	}
	result := &models.ConvertedFile{
		Data:        data,
		Name:        "output.txt",
		ContentType: "text/plain",
	}
	return result, nil

}
