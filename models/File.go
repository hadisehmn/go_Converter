package models

type File struct {
	// ID     int    `json:"id"`
	Name   string `json:"name"`
	Format string `json:"format"`
	Size   int64  `json:"size"`
	Data   []byte `json:"-"`
	// Path   string `json:"-"`

}
