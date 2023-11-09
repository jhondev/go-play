package fs

import (
	"bufio"
	"errors"
	"math"
	"os"
)

type File struct {
	blockSize int
}

func NewFile(blockSize int) File {
	return File{
		blockSize: blockSize,
	}
}

// Chunks calculates chunks from a file size
func (o File) Chunks(fsize int64) int {
	return int(math.Ceil(float64(fsize) / float64(o.blockSize)))
}

// Open opens a file split into at least two chunks.
func (f File) Open(name string) (*bufio.Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	finfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fsize := finfo.Size()
	fchunks := f.Chunks(fsize)
	// check required chunks
	if fchunks < 2 {
		return nil, errors.New("at least two chunks are required")
	}

	return bufio.NewReader(file), nil
}
