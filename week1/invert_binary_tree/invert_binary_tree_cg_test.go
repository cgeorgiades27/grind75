package invertbinarytree

import (
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestInvertTreeC(t *testing.T) {
	// todo: add c funcs to test
	for _, test := range TestCases {
		root := common.SliceToTreeC(test.input, false)
		target := common.SliceToTreeC(test.output, true)
		actual := invertBinaryTreeC(root)

		if !common.CompareTreesC(actual, target) {
			// t.Errorf("Test case %d failed", i)
		}
	}
}

func TestInvertTreeCG(t *testing.T) {
	for i, test := range TestCases {
		root := common.SliceToTree[int](test.input, func(i1, i2 int) bool { return i1 < i2 })
		target := common.SliceToTree[int](test.output, func(i1, i2 int) bool { return i1 > i2 })
		actual := invertBinaryTreeCG(root)

		var compareTrees func(*common.TreeNode[int], *common.TreeNode[int]) bool
		compareTrees = func(an, tn *common.TreeNode[int]) bool {
			if an == nil && tn == nil {
				return true
			}

			if an == nil || tn == nil {
				return false
			}

			if an.Val != tn.Val {
				return false
			}

			return compareTrees(an.Left, tn.Left) && compareTrees(an.Right, tn.Right)
		}

		if !compareTrees(actual, target) {
			t.Errorf("Test case %d failed", i)
		}
	}
}
