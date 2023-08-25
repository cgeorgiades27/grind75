package firstbadversion

import (
	"testing"
)

func firstBadVersionCg(n int, isBadVersion func(n int) bool) int {
	if isBadVersion(1) {
		return 1
	}

	high, low := n, 1
	for low <= high {
		mid := (high-low)/2 + low
		if isBadVersion(mid) &&
			!isBadVersion(mid-1) {
			return mid
		}

		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	return -1
}

func TestFirstBadVersionCg(t *testing.T) {
	for i, test := range TestCases {
		isBadVersion := func(n int) bool {
			return n >= test.output
		}
		output := firstBadVersionCg(test.input, isBadVersion)
		if output != test.output {
			t.Errorf("test(%d) expected %d, got %d", i, test.output, output)
		}
	}
}
