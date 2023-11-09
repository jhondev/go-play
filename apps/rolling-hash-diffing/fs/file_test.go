package fs_test

import (
	"testing"

	"github.com/jhondev/rolling-hash-diffing/fs"
)

func TestOpenFile(t *testing.T) {
	file1Path := "../testfiles/scores1.txt"
	blockSize := 16
	file := fs.NewFile(blockSize)
	_, err := file.Open(file1Path)
	if err != nil {
		t.Errorf("error opening file %s", file1Path)
	}
}

func TestChunksLen(t *testing.T) {
	file1Path := "../testfiles/scores1.txt"
	blockSize := 16
	file := fs.NewFile(blockSize)
	file1, err := file.Open(file1Path)
	if err != nil {
		t.Fatalf("error opening file %s", file1Path)
	}
	chunks := file.Chunks(int64(file1.Size()))
	if chunks != 256 {
		t.Errorf("expecting 256 chunks, got %d", chunks)
	}
}
