package maxdepthbinarytree

import (
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestMaxDepthBinaryTreeCg(t *testing.T) {
	for i, test := range TestCases {
		actual := maxDepthBinaryTreeCg(common.SliceToSearchTree(test.input, func(a, b int) bool { return a < b }, func(i1, i2 int) bool { return i1 == i2 }))
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}

func TestMaxDepthBinaryTreeC(t *testing.T) {
	for i, test := range TestCases {
		actual := maxDepthBinaryTreeC(common.SliceToSearchTreeC(test.input))
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}
