package balancedbinarytree

import (
	"github.com/cgeorgiades27/grind75/common"
)

/*
	#include <stdlib.h>
	#include <stdbool.h>

	#define MAX(a,b) ( (a)>(b) ? (a) : (b) )

	typedef struct TreeNode
	{
		struct TreeNode* left;
		struct TreeNode* right;
		int val;

	} TreeNode;

	typedef struct Result
	{
		bool is_balanced;
		int height;
	} Result;

	Result balanced_btree(TreeNode*);
	int btree_height(TreeNode*);

	int btree_height(TreeNode* root)
	{
		if (!root)
			return 0;

		return MAX(btree_height(root->left),btree_height(root->right)) + 1;
	}

	Result balanced_btree(TreeNode* root)
	{
		if (!root)
			return (Result){true,0};

		Result left = balanced_btree(root->left);
		if (!left.is_balanced)
			return (Result){false, left.height};

		Result right = balanced_btree(root->right);
		if (!right.is_balanced)
			return (Result){false, right.height};

		int res = left.height - right.height;
		if (res < 0)
			res *= -1;

		if (res > 1)
			return (Result){false, res};

		return (Result){true, MAX(left.height, right.height)+1};
	}
*/
import "C"

func balancedBinaryTreeCg(root *common.TreeNode[int]) bool {
	var recurser func(*common.TreeNode[int]) int
	recurser = func(root *common.TreeNode[int]) int {
		if root == nil {
			return 0
		}
		return max(recurser(root.Left), recurser(root.Right)) + 1
	}

	if root == nil {
		return true
	}

	var lheight int
	if root.Left != nil {
		lheight = recurser(root.Left)
	}

	var rheight int
	if root.Right != nil {
		rheight = recurser(root.Right)
	}

	return max(lheight, rheight)-min(lheight, rheight) < 2
}

func balancedBinaryTreeCg2(root *common.TreeNode[int]) bool {
	var balanced func(*common.TreeNode[int]) (bool, int)
	balanced = func(r *common.TreeNode[int]) (bool, int) {
		if r == nil {
			return true, 0
		}

		lBal, lHeight := balanced(r.Left)
		if !lBal {
			return false, lHeight
		}

		rBal, rHeight := balanced(r.Right)
		if !rBal {
			return false, rHeight
		}

		if absv(lHeight-rHeight) > 1 {
			return false, 0
		}

		return true, max(lHeight, rHeight) + 1
	}

	isBalanced, _ := balanced(root)
	return isBalanced
}

func absv(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
