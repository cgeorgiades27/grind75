package threesum

import (
	"slices"
	"testing"
)

func TestThreeSum(t *testing.T) {
	for i, test := range TestCases {
		actual := threeSum(test.Input)
		if len(actual) != len(test.Expected) {
			t.Fatalf("Test %d: Expected %v, got %v", i, test.Expected, actual)
		}
		for j, row := range test.Expected {
			if len(actual[j]) != len(row) {
				t.Fatalf("Test %d: Expected %v, got %v", i, row, actual[j])
			}
			slices.Sort(row)
			slices.Sort(actual[j])
			for k, e := range row {
				if actual[j][k] != e {
					t.Fatalf("Test %d: Expected %v, got %v", i, row, actual[j])
				}
			}
		}

	}
}
