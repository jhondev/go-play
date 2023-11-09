package diff

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

var Data1 = []byte("argentina 3 - 3 france")

func SignatureMock(blockSize int, t *testing.T) (*Signature, Diff) {
	diff := New(blockSize)
	file1 := bufio.NewReader(bytes.NewReader(Data1))
	signature, err := diff.Signature(file1)
	if err != nil {
		t.Fatalf("Error creating signature: %v", err)
	}
	return signature, diff
}

func TestCorrectSignature(t *testing.T) {
	signature, _ := SignatureMock(16, t)
	chunk1 := []byte("argentina 3 - 3 ")
	chunk2 := []byte("franceina 3 - 3 ")
	expected := &Signature{
		Checksums: []Checksum{
			{Weak: WeakChecksum(chunk1), Strong: StrongChecksum(chunk1)},
			{Weak: WeakChecksum(chunk2), Strong: StrongChecksum(chunk2)},
		},
	}
	if !reflect.DeepEqual(signature, expected) {
		t.Errorf("Expected signature %v\nGot %v", expected, signature)
	}
}
