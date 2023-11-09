package stats

import "math"

// sqrt ((Σ ((x - mean)^2) / n-1)
// mean can be calculated but for our case was already calculated
func StandardDeviation(data []float64, size int, mean *float64) float64 {
	variance := 0.0
	for _, x := range data {
		variance += (x - *mean) * (x - *mean)
	}
	variance = variance / float64(size)
	return math.Sqrt(variance)
}
