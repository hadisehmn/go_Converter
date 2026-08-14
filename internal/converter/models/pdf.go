package models

type PDF struct {
	File

	Encrypted bool `json:"encrypted"`
}
