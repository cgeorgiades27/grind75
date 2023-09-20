package diameterbinarytree

import (
	"unsafe"

	"github.com/cgeorgiades27/grind75/common"
)

/*
#include <stdlib.h>

typedef struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
} TreeNode;

int depth(int* m, TreeNode* r)
	{
		if (!r)
			return 0;

		int left = depth(m, r->left);
		int right = depth(m, r->right);

		if ((left + right) > *m)
			*m = left + right;

		return right > left ? right + 1 : left + 1;
}

int diameter_of_binary_tree_c(TreeNode* root)
{
	if (!root)
		return 0;

	int m = 0;
	depth(&m, root);

	return m;
}
*/
import "C"

func diameterOfBinaryTreeC(root *common.TreeNode[int]) int {
	arg := (*C.TreeNode)(unsafe.Pointer(root))
	res := C.diameter_of_binary_tree_c(arg)
	return (int)(res)
}

func diameterOfBinaryTreeCg(root *common.TreeNode[int]) int {
	if root == nil {
		return 0
	}

	var depth func(*int, *common.TreeNode[int]) int
	depth = func(m *int, r *common.TreeNode[int]) int {
		if r == nil {
			return 0
		}

		left := depth(m, r.Left)
		right := depth(m, r.Right)

		if left+right > *m {
			*m = left + right
		}

		if right > left {
			return right + 1
		}

		return left + 1
	}

	var m int
	depth(&m, root)

	return m
}
