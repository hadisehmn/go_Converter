package models

type PDF struct {
	File

	// PageCount int  `json:"page_count"`
	Encrypted bool `json:"encrypted"`
}
