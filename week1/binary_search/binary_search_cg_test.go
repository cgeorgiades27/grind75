package binarysearch

import "testing"

func TestBinarySearchCG(t *testing.T) {
	for i, test := range TestCases {
		actual := BinarySearchCG(test.nums, test.target)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}
