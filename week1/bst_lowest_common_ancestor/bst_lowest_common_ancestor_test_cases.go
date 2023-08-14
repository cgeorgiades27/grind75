package bstlowestcommonancestor

import "github.com/cgeorgiades27/grind75/common"

var (
	root1 *common.TreeNode[int]
	root2 *common.TreeNode[int]
	root3 *common.TreeNode[int]
	slc1  = []int{6, 2, 0, 4, 3, 5, 8, 7, 9}
	slc2  = []int{6, 2, 0, 4, 3, 5, 8, 7, 9}
	slc3  = []int{2, 1}
)

func init() {
	root1 = common.SliceToSearchTree[int](slc1, func(i1, i2 int) bool { return i1 < i2 }, func(i1, i2 int) bool { return i1 == i2 })
	root2 = common.SliceToSearchTree[int](slc2, func(i1, i2 int) bool { return i1 < i2 }, func(i1, i2 int) bool { return i1 == i2 })
	root3 = common.SliceToSearchTree[int](slc3, func(i1, i2 int) bool { return i1 < i2 }, func(i1, i2 int) bool { return i1 == i2 })
}

func findNode(i int, root *common.TreeNode[int]) *common.TreeNode[int] {
	if root == nil || root.Val == i {
		return root
	}

	left := findNode(i, root.Left)
	right := findNode(i, root.Right)

	if left != nil {
		return left
	}

	return right
}

var TestCases = []struct {
	root     *common.TreeNode[int]
	p        *common.TreeNode[int]
	q        *common.TreeNode[int]
	expected *common.TreeNode[int]
}{
	{
		root:     root1,
		p:        findNode(2, root1),
		q:        findNode(8, root1),
		expected: findNode(6, root1),
	},
	{
		root:     root2,
		p:        findNode(2, root2),
		q:        findNode(2, root2),
		expected: findNode(2, root2),
	},
	{
		root:     root3,
		p:        findNode(2, root3),
		q:        findNode(1, root3),
		expected: findNode(2, root3),
	},
}
