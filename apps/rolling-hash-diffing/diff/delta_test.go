package diff

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

var Data2 = []byte("croatia 2 - 1 morocco\nargentina 3 - 3 france")

func TestIndexesMapCreation(t *testing.T) {
	signature, _ := SignatureMock(8, t)
	indexes := WeakIndexes(signature)
	for i, check := range signature.Checksums {
		weakcheck := check.Weak
		strongcheck := check.Strong
		if indexes[weakcheck][strongcheck] != i {
			t.Errorf("Expected index %d for %d weak and %s strong", i, weakcheck, strongcheck)
		}
	}
}

func TestLookupIndex(t *testing.T) {
	signature, _ := SignatureMock(8, t)
	indexes := WeakIndexes(signature)
	windowtxt := "a 3 - 3 "
	window := []byte(windowtxt)
	weakcheck := WeakChecksum(window)
	index := Lookup(window, indexes, weakcheck)
	if index != 1 {
		t.Errorf("For window '%s' should get index 1", windowtxt)
	}
}

func TestDeltaChanges(t *testing.T) {
	signature, diff := SignatureMock(16, t)
	file2 := bufio.NewReader(bytes.NewReader(Data2))
	delta, err := diff.Delta(signature, file2)
	if err != nil {
		t.Errorf("Error computing delta: %v", err)
	}
	if len(delta) != 1 {
		t.Errorf("Delta should have 1 difference, got %d", len(delta))
	}
	difftxt := string(delta[0].Data)
	if difftxt != "croatia 2 - 1 morocco\n" {
		t.Errorf("Delta difference should be 'croatia 2 - 1 morocco', got '%s'", difftxt)
	}
}

func TestDeltaShifted(t *testing.T) {
	signature, diff := SignatureMock(4, t)
	datasft := []byte("argentina 0 - 0 france")
	file2 := bufio.NewReader(bytes.NewReader(datasft))
	delta, err := diff.Delta(signature, file2)
	if err != nil {
		t.Errorf("Error computing delta: %v", err)
	}
	for i, diff := range delta {
		fmt.Printf("Difference %d: '%s'", i+1, string(diff.Data))
	}
	expected := "argentina 0 - 0 "
	shifttxt := string(delta[4].Data)
	if shifttxt != expected {
		t.Errorf("Expecting '%s' in key index 4. Got '%s'", expected, shifttxt)
	}
}
