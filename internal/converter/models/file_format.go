package models

type FileFormat string

const (
	FormatPNG  FileFormat = "PNG"
	FormatJPG  FileFormat = "JPG"
	FormatJPEG FileFormat = "JPEG"
	FormatGIF  FileFormat = "GIF"
	FormatWEBP FileFormat = "WEBP"
	FormatPDF  FileFormat = "PDF"
	FormatICO  FileFormat = "ICO"
	FormatTXT  FileFormat = "TXT"
)
