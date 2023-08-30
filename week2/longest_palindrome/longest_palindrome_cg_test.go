package longestpalindrome

import (
	"testing"
)

func longestPalindromeCg(s string) int {
	slc := make([]int, 128)
	for _, r := range s {
		slc[r]++
	}

	count := 0
	hasSingleton := false
	for _, v := range slc {
		count += v
		if v%2 != 0 {
			hasSingleton = true
			count -= 1
		}
	}

	if hasSingleton {
		return count + 1
	}

	return count
}

func TestLongestPalindromeCg(t *testing.T) {
	for i, test := range TestCases {
		actual := longestPalindromeCg(test.input)
		if actual != test.output {
			t.Errorf("Test case %d: expected %d, got %d", i, test.output, actual)
		}
	}
}
