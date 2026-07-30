package service

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"

	"go-practice/CONVERTER/models"
)

type ImageService struct{}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (s *ImageService) ConvertIMG(file multipart.File, header *multipart.FileHeader, target string) (*models.ConvertedFile, error) {

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	img, err := imaging.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("invalid image")
	}

	var output bytes.Buffer
	var format imaging.Format

	switch strings.ToLower(target) {
	case "jpg", "jpeg":
		format = imaging.JPEG
	case "png":
		format = imaging.PNG
	default:
		return nil, errors.New("unsupported image format")
	}
	err = imaging.Encode(&output, img, format)
	if err != nil {
		return nil, err
	}

	ext := strings.TrimPrefix(filepath.Ext(header.Filename), ".")
	newName := strings.TrimSuffix(header.Filename, "."+ext) + "." + target

	return &models.ConvertedFile{
		Data:        output.Bytes(),
		Name:        newName,
		ContentType: "image/" + target,
	}, nil
}
