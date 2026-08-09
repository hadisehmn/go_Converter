package pdfservice

import (
	"bytes"
	"fmt"
	"go-practice/CONVERTER/models"
	"go-practice/CONVERTER/services/common"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/gen2brain/go-fitz"
)

func (s *PDFService) convertToImage(req models.ConvertRequest) (*models.ConvertedFile, error) {

	tempFile, err := os.CreateTemp("", "input-*.pdf")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	_, err = io.Copy(tempFile, req.File)
	if err != nil {
		return nil, err
	}

	doc, err := fitz.New(tempFile.Name())
	if err != nil {
		return nil, err
	}
	defer doc.Close()

	pageCount := doc.NumPage()

	if req.Page < 1 || req.Page > pageCount {
		return nil, fmt.Errorf("invalid page number: %d", req.Page)
	}

	pageIndex := req.Page - 1

	img, err := doc.Image(pageIndex)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	switch req.OutputFormat {
	case "png":
		err = png.Encode(&buffer, img)

	case "jpg", "jpeg":
		err = jpeg.Encode(&buffer, img, &jpeg.Options{
			Quality: 90,
		})
	}

	if err != nil {
		return nil, err
	}

	contentType := "image/png"

	if req.OutputFormat == "jpg" || req.OutputFormat == "jpeg" {
		contentType = "image/jpeg"
	}

	result := common.BuildConvertedFile(
		buffer,
		req,
		contentType,
	)

	return result, nil

}
