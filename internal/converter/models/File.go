package models

type File struct {
	Name   string     `json:"name"`
	Format FileFormat `json:"format"`
	Size   int64      `json:"size"`
	Data   []byte     `json:"-"`
}
