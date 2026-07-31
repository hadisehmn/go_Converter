package controller

import (
	"encoding/json"
	"go-practice/CONVERTER/models"
	service "go-practice/CONVERTER/services"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ConvertController struct {
	service *service.ConvertService
}

func NewConvertController(service *service.ConvertService) *ConvertController {
	return &ConvertController{
		service: service,
	}
}

func (h *ConvertController) ConvertFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
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

	ext := strings.TrimPrefix(filepath.Ext(header.Filename), ".")

	convertReq := models.ConvertRequest{
		File:         file,
		Header:       header,
		InputFormat:  strings.ToLower(ext),
		OutputFormat: strings.ToLower(targetFormat),
	}
	convertedFile, err := h.service.Convert(convertReq)
	if err != nil {
		http.Error(w, "conversion failed", 500)
		return
	}
	err = os.MkdirAll("downloads", 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(
		filepath.Join("downloads", convertedFile.Name),
		convertedFile.Data,
		0644,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]string{
		"message":      "File converted successfully",
		"download_url": "/downloads/" + convertedFile.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
