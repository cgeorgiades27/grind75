package balancedbinarytree

import "github.com/cgeorgiades27/grind75/common"

var TestCases = []struct {
	input  *common.TreeNode[int]
	output bool
}{
	{
		input: &common.TreeNode[int]{
			Val:  1,
			Left: &common.TreeNode[int]{Val: 9},
			Right: &common.TreeNode[int]{
				Val:   20,
				Right: &common.TreeNode[int]{Val: 7},
				Left:  &common.TreeNode[int]{Val: 15},
			},
		},
		output: true,
	},
	{
		input: &common.TreeNode[int]{
			Val:   1,
			Right: &common.TreeNode[int]{Val: 2},
			Left: &common.TreeNode[int]{
				Val:   2,
				Right: &common.TreeNode[int]{Val: 3},
				Left: &common.TreeNode[int]{
					Val:   3,
					Right: &common.TreeNode[int]{Val: 4},
					Left:  &common.TreeNode[int]{Val: 4},
				},
			},
		},
		output: false,
	},
	{
		input:  nil,
		output: true,
	},
}
