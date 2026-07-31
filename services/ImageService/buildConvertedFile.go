package imageservice

import (
	"bytes"
	"go-practice/CONVERTER/models"
	"path/filepath"
	"strings"
)

func (s *ImageService) buildConvertedFile(output bytes.Buffer, req models.ConvertRequest, contentType string) *models.ConvertedFile {

	ext := strings.TrimPrefix(filepath.Ext(req.Header.Filename), ".")
	newName := strings.TrimSuffix(req.Header.Filename, "."+ext) + "." + req.OutputFormat

	return &models.ConvertedFile{
		Data:        output.Bytes(),
		Name:        newName,
		ContentType: contentType,
	}
}
