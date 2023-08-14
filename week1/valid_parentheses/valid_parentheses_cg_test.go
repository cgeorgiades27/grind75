package validparentheses

import (
	"testing"
)

// validParenthesesCG is a solution for https://leetcode.com/problems/valid-parentheses/
func validParenthesesCG(s string) bool {
	comps := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, b := range s {
		if b != ')' &&
			b != ']' &&
			b != '}' {
			stack = append(stack, b)
			continue
		}

		if len(stack) == 0 || comps[b] != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

// validParenthesesNoSlicingCG is a solution for https://leetcode.com/problems/valid-parentheses/
func validParenthesesNoSlicingCG(s string) bool {
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

		if sPtr == 0 {
			return false
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
		actual := validParenthesesCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, test.output, actual)
		}
	}
}

func TestValidParenthesesNoSlicingCG(t *testing.T) {
	for i, test := range TestCases {
		actual := validParenthesesNoSlicingCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, test.output, actual)
		}
	}
}
