package maxdepthbinarytree

import (
	"unsafe"

	"github.com/cgeorgiades27/grind75/common"
)

/*
#include <stdio.h>

typedef struct TreeNode
{
	int val;
	struct TreeNode* left;
	struct TreeNode* right;
} TreeNode;

int max_depth_binary_tree_cg(TreeNode* root)
{
	if (!root)
		return 0;

	int left = max_depth_binary_tree_cg(root->left);
	int right = max_depth_binary_tree_cg(root->right);

	return left > right ? ++left : ++right;
}
*/
import "C"

func maxDepthBinaryTreeCg(root *common.TreeNode[int]) int {
	if root == nil {
		return 0
	}

	return max(
		maxDepthBinaryTreeCg(root.Left),
		maxDepthBinaryTreeCg(root.Right),
	) + 1
}

func maxDepthBinaryTreeC(root *common.TreeNode[int]) int {
	arg := (*C.TreeNode)(unsafe.Pointer(root))
	return (int)(C.max_depth_binary_tree_cg(arg))
}
