package models

type File struct {
	Name   string `json:"name"`
	Format string `json:"format"`
	Size   int64  `json:"size"`
	Data   []byte `json:"-"`
}
