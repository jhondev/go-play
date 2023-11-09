package processor_test

import (
	"testing"

	"logprocessor.37Widgets.co/processor"
)

func TestPrecisionRule(t *testing.T) {
	tt := []struct {
		desc     string
		total    float64
		counter  int
		readings []float64
		ref      float64
		expected string
	}{
		{desc: "thermometer temp-1",
			total:    905.2,
			counter:  13,
			readings: []float64{72.4, 76.0, 79.1, 75.6, 71.2, 71.4, 69.2, 65.2, 62.8, 61.4, 64.0, 67.5, 69.4},
			ref:      70.0,
			expected: processor.Precise},
		{desc: "thermometer temp-2",
			total:    352.2,
			counter:  5,
			readings: []float64{69.5, 70.1, 71.3, 71.5, 69.8},
			ref:      70.0,
			expected: processor.UltraPrecise},
	}

	for _, tc := range tt {
		result := processor.PrecisionRule(tc.total, tc.counter, tc.readings, tc.ref)
		t.Log(tc.desc)
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}

func TestDiscardRule(t *testing.T) {
	tt := []struct {
		desc     string
		total    float64
		counter  int
		readings []float64
		ref      float64
		expected string
	}{
		{desc: "humidity hum-1",
			total:    135.6,
			counter:  3,
			readings: []float64{45.2, 45.3, 45.1},
			ref:      45.0,
			expected: processor.OK},
		{desc: "humidity hum-2",
			total:    219.1,
			counter:  5,
			readings: []float64{44.4, 43.9, 44.9, 43.8, 42.1},
			ref:      45.0,
			expected: processor.Discard},
	}

	for _, tc := range tt {
		result := processor.DiscardRule(tc.total, tc.counter, tc.readings, tc.ref)
		t.Log(tc.desc)
		if result != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, result)
		}
	}
}
