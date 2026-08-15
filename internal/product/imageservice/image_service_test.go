package imageservice

import (
	"go-practice/CONVERTER/internal/converter/models"
	"image/png"
	"mime/multipart"
	"os"
	"testing"
)

func TestConvertToGIF(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatGIF,
		Page:         1,
	}

	result, err := service.convertToGIF(img, req)
	if err != nil {
		t.Fatalf("convertToGIF failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("GIF generated successfully: %+v", result)

}

func TestConvertToICO(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatICO,
		Page:         1,
	}

	result, err := service.convertToICO(img, req)
	if err != nil {
		t.Fatalf("convertToICOfailed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("ICO generated successfully: %+v", result)

}

func TestConvertToJPEG(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatJPEG,
		Page:         1,
	}

	result, err := service.convertToJPEG(img, req)
	if err != nil {
		t.Fatalf("convertToJPEGfailed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("JPEG generated successfully: %+v", result)

}

func TestConvertToPDF(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatPDF,
		Page:         1,
	}

	result, err := service.convertToPDF(img, req)
	if err != nil {
		t.Fatalf("convertToPDFfailed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("PDF generated successfully: %+v", result)

}

func TestConvertToPNG(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatPNG,
		Page:         1,
	}

	result, err := service.convertToPNG(img, req)
	if err != nil {
		t.Fatalf("convertToPNGfailed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("PNG generated successfully: %+v", result)

}

func TestConvertToWEBP(t *testing.T) {
	service := &ImageService{}

	file, err := os.Open("testdata/input.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	req := models.ConvertRequest{
		Header: &multipart.FileHeader{
			Filename: "input.png",
		},
		InputFormat:  models.FormatPNG,
		OutputFormat: models.FormatWEBP,
		Page:         1,
	}

	result, err := service.convertToWEBP(img, req)
	if err != nil {
		t.Fatalf("convertToWEBPfailed: %v", err)
	}

	if result == nil {
		t.Fatal("expected result, got nil")
	}

	t.Logf("WEBP generated successfully: %+v", result)

}
