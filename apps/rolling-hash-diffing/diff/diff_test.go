package diff

import (
	"testing"
)

func TestCompareWithFiles(t *testing.T) {
	file1Path := "../testfiles/scores1.txt"
	file2Path := "../testfiles/scores2.txt"
	blockSize := 16

	diff := New(blockSize)
	delta, err := diff.Compare(file1Path, file2Path)
	if err != nil {
		t.Fatalf("error comparing files: %v", err)
	}

	difftxt := string(delta[0].Data)
	if difftxt != "croatia 2 - 1 morocco\n" {
		t.Errorf("Delta difference should be 'croatia 2 - 1 morocco', got '%s'", difftxt)
	}
}
