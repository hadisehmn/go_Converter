package service

import (
	"errors"
	"go-practice/CONVERTER/models"
	imageservice "go-practice/CONVERTER/services/ImageService"
	"strings"
)

type ConvertService struct {
	imageService *imageservice.ImageService
}

func NewConvertService() *ConvertService {
	return &ConvertService{
		imageService: imageservice.NewImageService(),
	}
}
func isImage(format string) bool {
	switch strings.ToLower(format) {
	case "jpg", "jpeg", "png", "gif", "bmp", "webp", "tiff", "avif", "heic":
		return true
	default:
		return false
	}
}

func (s *ConvertService) Convert(req models.ConvertRequest) (*models.ConvertedFile, error) {

	switch {

	case isImage(req.InputFormat):
		return s.imageService.ConvertIMG(req)

	default:
		return nil, errors.New("unsupported input format")
	}
}
