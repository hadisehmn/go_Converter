package imageservice

import (
	"bytes"
	"image"

	"go-practice/CONVERTER/models"

	"github.com/disintegration/imaging"
	"github.com/phpdave11/gofpdf"
)

func (s *ImageService) convertToPDF(img image.Image, req models.ConvertRequest) (*models.ConvertedFile, error) {

	var imageBuffer bytes.Buffer

	err := imaging.Encode(&imageBuffer, img, imaging.PNG)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	opt := gofpdf.ImageOptions{
		ImageType: "PNG",
	}

	pdf.RegisterImageOptionsReader("image", opt, &imageBuffer)

	pdf.ImageOptions("image", 10, 10, 190, 0, false, opt, 0, "")

	var pdfBuffer bytes.Buffer

	err = pdf.Output(&pdfBuffer)
	if err != nil {
		return nil, err
	}

	return s.buildConvertedFile(pdfBuffer, req, "application/pdf"), nil
}
