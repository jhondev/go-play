package stats_test

import (
	"math"
	"testing"

	"logprocessor.37Widgets.co/stats"
)

func TestStandardVariation(t *testing.T) {
	// Can be improved with table tests
	sample := []float64{2, 1, 3, 2, 4}
	mean := 2.4
	sdev := math.Round(stats.StandardDeviation(sample, 5, &mean)*100) / 100

	if sdev != 1.02 {
		t.Errorf("Standard Deviation Incorrect: %f", sdev)
	}
}
