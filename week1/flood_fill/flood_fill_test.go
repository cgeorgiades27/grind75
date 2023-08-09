package floodfill

import (
	"slices"
	"testing"
)

func TestFloodFillCG(t *testing.T) {
	for i, test := range TestCases {
		actual := FloodFillCG(test.image, test.sr, test.sc, test.color)
		for j, slc := range test.output {
			if !slices.Equal(slc, actual[j]) {
				t.Errorf("Test %d: expected %v, actual %v", i, slc, actual[j])
			}
		}
	}
}
