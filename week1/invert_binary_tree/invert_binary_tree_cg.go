package invertbinarytree

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

TreeNode *invert_binary_tree(struct TreeNode *root)
{
    if (!root)
        return root;

    struct TreeNode *temp = root->left;
    root->left = root->right;
    root->right = temp;

    root->left = invert_binary_tree(root->left);
    root->right = invert_binary_tree(root->right);

    return root;
}
*/
import "C"

func InvertBinaryTreeC(root *common.TreeNode[int]) *common.TreeNode[int] {
	arg := (*C.TreeNode)(unsafe.Pointer(root))
	res := C.invert_binary_tree(arg)
	return (*common.TreeNode[int])(unsafe.Pointer(res))
}

func InvertBinaryTreeCG(root *common.TreeNode[int]) *common.TreeNode[int] {
	if root == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	root.Left = InvertBinaryTreeC(root.Left)
	root.Right = InvertBinaryTreeC(root.Right)

	return root
}
