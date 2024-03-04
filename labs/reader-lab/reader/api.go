package reader

import (
	"io"
)

type Url string

func NewAPI(url Url) io.Reader {
	return API{}
}

type API struct{}

// Read implements io.Reader.
func (f API) Read(p []byte) (n int, err error) {
	return 0, nil
}
