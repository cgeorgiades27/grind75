package floodfill

import (
	"slices"
	"testing"
)

func TestFloodFillCG(t *testing.T) {
	for i, test := range TestCases {
		actual := floodFillCG(test.image, test.sr, test.sc, test.color)
		for j, slc := range test.output {
			if !slices.Equal(slc, actual[j]) {
				t.Errorf("Test %d: expected %v, actual %v", i, slc, actual[j])
			}
		}
	}
}

func TestFloodFillCR(t *testing.T) {
	for i, test := range TestCases {
		actual := floodFillCR(test.image, test.sr, test.sc, test.color)
		for j, slc := range test.output {
			if !slices.Equal(slc, actual[j]) {
				t.Errorf("Test %d: expected %v, actual %v", i, slc, actual[j])
			}
		}
	}
}

func BenchmarkFloodFillCR(b *testing.B) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	color := 2
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		floodFillCR(image, sr, sc, color)
	}
}
