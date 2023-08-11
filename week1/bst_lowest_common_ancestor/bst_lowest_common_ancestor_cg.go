package bstlowestcommonancestor

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

TreeNode *lowest_common_ancestor(TreeNode *root, TreeNode *p, TreeNode *q)
{
	if (!root)
		return root;

	if (root == p  || root == q)
		return root;

	TreeNode *left = lowest_common_ancestor(root->left, p, q);
	TreeNode *right = lowest_common_ancestor(root->right, p, q);

	if (left && right)
		return root;

	if (left)
		return left;

	return right;
}
*/
import "C"

func lowestCommonAncestorCGC(root, p, q *common.TreeNode[int]) *common.TreeNode[int] {
	rc := (*C.TreeNode)(unsafe.Pointer(root))
	pc := (*C.TreeNode)(unsafe.Pointer(p))
	qc := (*C.TreeNode)(unsafe.Pointer(q))
	res := C.lowest_common_ancestor(rc, pc, qc)
	return (*common.TreeNode[int])(unsafe.Pointer(res))
}

func lowestCommonAncestorCG(root, p, q *common.TreeNode[int]) *common.TreeNode[int] {
	if root == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestorCG(root.Left, p, q)
	right := lowestCommonAncestorCG(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
