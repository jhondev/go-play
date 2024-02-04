package main

import (
	"testing"

	"github.com/spatialcurrent/go-math/pkg/math"
)

func src() []int {
	src := []int{}
	for i := 0; i < 100000000; i++ {
		src = append(src, i)
	}
	return src
}

func TestSlns(t *testing.T) {
	tt := []struct {
		name string
		fn   func([]int) int
	}{{
		name: "normal",
		fn:   sum,
	}, {
		name: "optimized",
		fn:   sumopt,
	}}
	src := src()
	expected, _ := math.Sum(src)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := tc.fn(src)
			if r != expected {
				t.Fatalf("Expected %d got %d", expected, r)
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	src := src()
	sum(src)
	// t := sum(src)
	// b.Log(t)
}

func BenchmarkSumopt(b *testing.B) {
	src := src()
	sumopt(src)
	// t := sumopt(src)
	// b.Log(t)
}
