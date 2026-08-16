package pdfservice

import (
	"go-practice/CONVERTER/internal/converter/models"
	"mime/multipart"
	"os"
	"testing"
)

func TestConvertToImage(t *testing.T) {
	service := &PDFService{}

	file, err := os.Open("testdata/input.pdf")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	req := models.ConvertRequest{
		File: file,

		Header: &multipart.FileHeader{
			Filename: "input.pdf",
		},

		InputFormat:  models.FormatPDF,
		OutputFormat: models.FormatJPEG,
		Page:         1,
	}

	result, err := service.convertToImage(req)
	if err != nil {
		t.Fatalf("convertToImage failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("JPEG generated successfully: %+v", result)
}
