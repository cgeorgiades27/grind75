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
			// todo - get this to work
		}
	}
}

func TestInvertTreeCG(t *testing.T) {
	for i, test := range TestCases {
		root := common.SliceToTree[int](test.input, func(i1, i2 int) bool { return i1 < i2 })
		target := common.SliceToTree[int](test.output, func(i1, i2 int) bool { return i1 > i2 })
		actual := invertBinaryTreeCG(root)

		var recurser func(*common.TreeNode[int], *common.TreeNode[int])
		recurser = func(an, tn *common.TreeNode[int]) {
			if an == nil && tn == nil {
				return
			}

			if an == nil || tn == nil {
				t.Errorf("Test %d: nil mismatch expected %v, got %v", i, tn, an)
			}

			if an.Val != tn.Val {
				t.Errorf("Test %d: val mismatch expected %v, got %v", i, tn.Val, an.Val)
			}

			recurser(an.Left, tn.Left)
			recurser(an.Right, tn.Right)
		}

		recurser(actual, target)
	}
}
