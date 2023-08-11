package balancedbinarytree

import "github.com/cgeorgiades27/grind75/common"

func BalancedBinaryTree(root *common.TreeNode[int]) bool {
	var recurser func(int, *common.TreeNode[int]) int
	recurser = func(h int, root *common.TreeNode[int]) int {
		if root == nil {
			return h
		}

		h++
		return recurser(h, root.Left) + recurser(h, root.Right)
	}
}
