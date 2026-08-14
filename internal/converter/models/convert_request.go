package models

import "mime/multipart"

type ConvertRequest struct {
	File   multipart.File
	Header *multipart.FileHeader

	InputFormat  FileFormat
	OutputFormat FileFormat

	Page int
}
