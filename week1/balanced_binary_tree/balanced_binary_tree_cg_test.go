package balancedbinarytree

import (
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func BalancedBinaryTree(root *common.TreeNode[int]) bool {
	var recurser func(*common.TreeNode[int]) int
	recurser = func(root *common.TreeNode[int]) int {
		if root == nil {
			return 0
		}
		return max(recurser(root.Left), recurser(root.Right)) + 1
	}

	if root == nil {
		return true
	}

	var lheight int
	if root.Left != nil {
		lheight = recurser(root.Left)
	}

	var rheight int
	if root.Right != nil {
		rheight = recurser(root.Right)
	}

	return max(lheight, rheight)-min(lheight, rheight) < 2
}

func TestBalancedBinaryTreeCG(t *testing.T) {
	for i, test := range TestCases {
		if BalancedBinaryTree(test.input) != test.output {
			t.Errorf("Test %d Failed", i)
		}
	}
}
