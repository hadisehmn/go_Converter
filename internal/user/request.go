package user

import (
	"go-practice/CONVERTER/internal/converter"
	"go-practice/CONVERTER/internal/converter/models"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type ConvertController struct {
	service *converter.ConvertService
}

func NewConvertController(service *converter.ConvertService) *ConvertController {
	return &ConvertController{
		service: service,
	}
}

func (h *ConvertController) ConvertFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file missing", http.StatusBadRequest)
		return
	}
	defer file.Close()

	targetFormat := r.FormValue("target")
	if targetFormat == "" {
		http.Error(w, "target format is required", http.StatusBadRequest)
		return
	}

	ext := strings.TrimPrefix(
		filepath.Ext(header.Filename),
		".",
	)

	page := 1

	if pageStr := r.FormValue("page"); pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			http.Error(w, "invalid page", http.StatusBadRequest)
			return
		}
	}

	inputFormat := models.FileFormat(strings.ToUpper(ext))
	outputFormat := models.FileFormat(strings.ToUpper(targetFormat))

	convertReq := models.ConvertRequest{
		File:         file,
		Header:       header,
		InputFormat:  inputFormat,
		OutputFormat: outputFormat,
		Page:         page,
	}

	convertedFile, err := h.service.Convert(convertReq)
	if err != nil {
		log.Println("CONVERSION ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", convertedFile.ContentType)

	w.Header().Set(
		"Content-Disposition",
		`attachment; filename="`+convertedFile.Name+`"`,
	)

	w.Header().Set(
		"Content-Length",
		strconv.Itoa(len(convertedFile.Data)),
	)

	_, err = w.Write(convertedFile.Data)
	if err != nil {
		log.Println("RESPONSE ERROR:", err)
		return
	}
}
