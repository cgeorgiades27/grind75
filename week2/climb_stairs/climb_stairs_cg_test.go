package climbstairs

import "testing"

func climbStairsCg(n int) int {
	if n < 3 {
		return n
	}

	return climbStairsCg(n-1) + climbStairsCg(n-2)
}

func TestClimbStairsCg(t *testing.T) {
	for i, test := range TestCases {
		actual := climbStairsCg(test.n)
		if actual != test.expected {
			t.Errorf("Test #%d: got %d, want %d", i+1, actual, test.expected)
		}
	}
}
