package models

type Image struct {
	File
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	OutputFormat string `json:"output_format"`
}
