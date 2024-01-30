package common

import "testing"

func TestCheckSilceEquality(t *testing.T) {
	for i, test := range []struct {
		s1       []int
		s2       []int
		expected bool
	}{
		{
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 3},
			expected: true,
		},
		{
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 4},
			expected: false,
		},
		{
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2},
			expected: false,
		},
	} {
		if CheckSliceEquality[int](test.s1, test.s2) != test.expected {
			t.Fatalf("Test %d: Expected %v, got %v", i, test.expected, !test.expected)
		}
	}
}
