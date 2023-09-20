package diameterbinarytree

import "github.com/cgeorgiades27/grind75/common"

var TestCases = []struct {
	input    *common.TreeNode[int]
	expected int
}{
	{
		input: &common.TreeNode[int]{
			Val: 1,
			Left: &common.TreeNode[int]{
				Val:   2,
				Left:  &common.TreeNode[int]{Val: 4},
				Right: &common.TreeNode[int]{Val: 5},
			},
			Right: &common.TreeNode[int]{Val: 3},
		},
		expected: 3,
	},
	{
		input:    &common.TreeNode[int]{Val: 1, Left: &common.TreeNode[int]{Val: 2}},
		expected: 1,
	},
}
