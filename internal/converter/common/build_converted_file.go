package common

import (
	"bytes"
	"path/filepath"
	"strings"

	"go-practice/CONVERTER/internal/converter/models"
)

func BuildConvertedFile(output bytes.Buffer, req models.ConvertRequest, contentType string) *models.ConvertedFile {

	baseName := strings.TrimSuffix(req.Header.Filename, filepath.Ext(req.Header.Filename))

	newName := baseName + "." + string(req.OutputFormat)
	return &models.ConvertedFile{
		Data:        output.Bytes(),
		Name:        newName,
		ContentType: contentType,
	}
}
