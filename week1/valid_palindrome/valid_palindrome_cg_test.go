package validpalindrome

import (
	"strings"
	"testing"
)

func isPalindromeCG(s string) bool {
	isAlnum := func(ch rune) bool {
		return (ch >= '0' && ch <= '9') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= 'a' && ch <= 'z')
	}

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

		if !strings.EqualFold(string(s[begin]), string(s[end])) {
			return false
		}

		begin++
		end--
	}

	return true
}

func TestIsPalindromeCG(t *testing.T) {
	for i, test := range TestCases {
		actual := isPalindromeCG(test.input)
		if actual != test.output {
			t.Errorf("Test case %d: expected %t, got %t", i, test.output, actual)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	for i, test := range []struct {
		c        rune
		expected bool
	}{
		{
			c:        'a',
			expected: true,
		},
		{
			c:        '1',
			expected: false,
		},
		{
			c:        '.',
			expected: false,
		},
		{
			c:        'Z',
			expected: true,
		},
		{
			c:        ' ',
			expected: false,
		},
	} {
		actual := isAlpha(test.c)
		if actual != test.expected {
			t.Errorf("test %d failed: expected %t, got: %t", i, test.expected, actual)
		}
	}
}

func TestValidPalindrome(t *testing.T) {
	for i, test := range TestCases {
		actual := validPalindrome(test.input)
		if actual != test.output {
			t.Errorf("Test case %d: expected %t, got %t", i, test.output, actual)
		}
	}
}

func TestIsIntPalindrome(t *testing.T) {
	for i, test := range TestCasesInt {
		actual := isIntPalindrome(test.input)
		if actual != test.expected {
			t.Errorf("Test case %d: expected %t, got %t", i, test.expected, actual)
		}
	}
}
