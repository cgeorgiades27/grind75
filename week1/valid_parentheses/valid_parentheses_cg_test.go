package validparentheses

import (
	"testing"
)

// ValidParenthesesCG is a solution for https://leetcode.com/problems/valid-parentheses/
func ValidParenthesesCG(s string) bool {
	var stack []rune
	comps := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, b := range s {
		if b != ')' &&
			b != ']' &&
			b != '}' {
			stack = append(stack, b)
			continue
		}

		if comps[b] != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// ValidParenthesesNoSlicingCG is a solution for https://leetcode.com/problems/valid-parentheses/
func ValidParenthesesNoSlicingCG(s string) bool {
	comps := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune, len(s))
	sPtr := 0

	for _, b := range s {
		if b != ')' &&
			b != ']' &&
			b != '}' {
			stack[sPtr] = b
			sPtr++
			continue
		}

		sPtr--
		if comps[b] != stack[sPtr] {
			return false
		}
	}

	return sPtr == 0
}

func TestValidParenthesesCG(t *testing.T) {
	for i, test := range TestCases {
		actual := ValidParenthesesCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, test.output, actual)
		}
	}
}

func TestValidParenthesesNoSlicingCG(t *testing.T) {
	for i, test := range TestCases {
		actual := ValidParenthesesNoSlicingCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, test.output, actual)
		}
	}
}
