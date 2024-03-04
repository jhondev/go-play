package util

import (
	"log"
	"testing"
)

func FatalIf(err error) {
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}

func FatalTestIf(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
