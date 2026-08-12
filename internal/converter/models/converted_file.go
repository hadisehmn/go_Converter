package models

type ConvertedFile struct {
	Data        []byte `json:"data"`
	Name        string `json:"name"`
	ContentType string `json:"content_type"`
}
