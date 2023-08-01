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

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "()[][[[",
			output: false,
		},
	}

	for i, test := range tests {
		actual := ValidParenthesesCG(test.input)
		if actual != test.output {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, test.output, actual)
		}
	}
}
