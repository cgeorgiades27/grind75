package validpalindrome

import (
	"strings"
	"testing"
)

func isPalindromeCG(s string) bool {
	begin, end := 0, len(s)-1
	for begin < end {
		for !(isAlnum(rune(s[begin]))) &&
			begin < end {
			begin++
		}

		for !(isAlnum(rune(s[end]))) &&
			end > begin {
			end--
		}

		if strings.ToLower(string(s[begin])) != strings.ToLower(string(s[end])) {
			return false
		}

		begin++
		end--
	}

	return true
}

func isAlnum(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func TestIsPalindromeCG(t *testing.T) {
	for i, test := range TestCases {
		actual := isPalindromeCG(test.input)
		if actual != test.output {
			t.Errorf("Test case %d: expected %t, got %t", i, test.output, actual)
		}
	}
}
