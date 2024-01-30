package binarysearch

import "testing"

func binarySearchCG(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func TestBinarySearchCG(t *testing.T) {
	for i, test := range TestCases {
		actual := binarySearchCG(test.nums, test.target)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}

func TestBinarySearchCg(t *testing.T) {
	for i, test := range TestCases {
		actual := binarySearchCg(test.target, test.nums)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}

func TestBinarySearchCg2(t *testing.T) {
	for i, test := range TestCases {
		actual := binarySearchCg2(test.target, test.nums)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}
