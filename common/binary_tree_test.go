package common

import "testing"

func TestTreeToSlice(t *testing.T) {
	var (
		tree1 *TreeNode[int]
		tree2 *TreeNode[int]
		tree3 *TreeNode[int]
	)

	tree1 = &TreeNode[int]{
		Val: 1,
		Left: &TreeNode[int]{
			Val: 2,
			Left: &TreeNode[int]{
				Val:  3,
				Left: &TreeNode[int]{Val: 4},
			},
		},
		Right: &TreeNode[int]{
			Val: 5,
			Right: &TreeNode[int]{
				Val:   6,
				Right: &TreeNode[int]{Val: 7},
			},
		},
	}

	tree2 = &TreeNode[int]{Val: 1}
	tree3 = nil

	tests := []struct {
		input  *TreeNode[int]
		output []int
	}{
		{
			input:  tree1,
			output: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			input:  tree2,
			output: []int{1},
		},
		{
			input:  tree3,
			output: []int{},
		},
	}

	for i, test := range tests {
		actual := TreeToSlice[int](test.input)
		if len(actual) != len(test.output) {
			t.Errorf("test %d failed, wanted len of: %d, got: %d", i, len(test.output), len(actual))
		}

		for j, elem := range test.output {
			if elem != actual[j] {
				t.Errorf("test %d failed, wanted: %d, got: %d", i, elem, actual[j])
			}
		}
	}

}

func TestSliceToTree(t *testing.T) {
	var (
		tree1 *TreeNode[int]
		tree2 *TreeNode[int]
		tree3 *TreeNode[int]
		slc1  = []int{1, 2, 3, 4, 5, 6, 7}
		slc2  = []int{1}
		slc3  = []int{}
	)

	for _, elem := range slc1 {
		tree1 = insertNode[int](tree1, elem, func(i1, i2 int) bool { return i1 < i2 })
	}
	tree2 = &TreeNode[int]{Val: 1}
	tree3 = nil

	tests := []struct {
		input  []int
		output *TreeNode[int]
	}{
		{
			input:  slc1,
			output: tree1,
		},
		{
			input:  slc2,
			output: tree2,
		},
		{
			input:  slc3,
			output: tree3,
		},
	}

	for i, test := range tests {
		actual := SliceToTree[int](test.input, func(i1, i2 int) bool { return i1 < i2 })

		var preorderCheck func(root1, root2 *TreeNode[int]) bool
		preorderCheck = func(root1, root2 *TreeNode[int]) bool {
			if root1 == nil && root2 == nil {
				return true
			}

			if root1 == nil {
				return false
			}

			if root2 == nil {
				return false
			}

			if root1.Val != root2.Val {
				return false
			}

			return preorderCheck(root1.Left, root2.Left) && preorderCheck(root1.Right, root2.Right)
		}

		res := preorderCheck(actual, test.output)

		if !res {
			t.Errorf("test %d failed, wanted: %t, got: %t", i, true, res)
		}
	}
}
