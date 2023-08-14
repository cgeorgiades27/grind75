package bstlowestcommonancestor

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	for i, test := range TestCases {
		actual := lowestCommonAncestorCG(test.root, test.p, test.q)
		if actual != test.expected {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.expected, actual)
		}
	}
}

func TestLowestCommonAncestorC(t *testing.T) {
	for i, test := range TestCases {
		actual := lowestCommonAncestorC(test.root, test.p, test.q)
		if actual != test.expected {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.expected, actual)
		}
	}
}
