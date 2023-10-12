package balancedbinarytree

import (
	"testing"
)

func TestBalancedBinaryTreeCG(t *testing.T) {
	for i, test := range TestCases {
		actual := balancedBinaryTreeCg(test.input)
		if actual != test.output {
			t.Errorf("test %d, wanted %v got %v", i, test.output, actual)
		}
	}
}

func TestBalancedBinaryTreeCg2(t *testing.T) {
	for i, test := range TestCases {
		actual := balancedBinaryTreeCg2(test.input)
		if actual != test.output {
			t.Errorf("test %d, wanted %v got %v", i, test.output, actual)
		}
	}
}
