package common

import (
	"slices"
)

/*
#include <stdlib.h>
#include <stdbool.h>

typedef struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;

} TreeNode;

typedef bool (*comparitor)(int, int);

int less(int a, int b) { return a < b; };
int greater(int a, int b) { return a > b; };

TreeNode *arr_to_tree(int *, comparitor);
TreeNode *insert_node(TreeNode *, int, comparitor);

TreeNode *arr_to_tree(int *arr, comparitor cmp)
{
    if (!arr)
        return NULL;

    TreeNode *root = malloc(sizeof(TreeNode));
    for (int i = 0; i < sizeof(arr) / sizeof(int); ++i)
        root = insert_node(root, arr[i], cmp);

    return root;
}

TreeNode *insert_node(TreeNode *node, int val, comparitor cmp)
{
    if (!node)
    {
        node = malloc(sizeof(TreeNode));
        node->val = val;
        return node;
    }

    if (cmp(val, node->val))
        node->left = insert_node(node->left, val, cmp);
    else
        node->right = insert_node(node->left, val, cmp);

    return node;
}
*/
import "C"

// Definition for a binary tree node.
type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func TreeToSlice[T any](root *TreeNode[T]) []T {
	var slc []T
	var preorderRecurser func(root *TreeNode[T])
	preorderRecurser = func(root *TreeNode[T]) {
		if root == nil {
			return
		}

		slc = append(slc, root.Val)
		preorderRecurser(root.Left)
		preorderRecurser(root.Right)
		return
	}

	preorderRecurser(root)
	return slc
}

func SliceToTree[T any](slc []T, less func(T, T) bool) *TreeNode[T] {
	if len(slc) == 0 {
		return nil
	}

	var root *TreeNode[T] = nil
	for _, elem := range slc {
		root = insertNode(root, elem, less)
	}

	return root
}

func SliceToSearchTree[T any](slc []T, less func(T, T) bool, equal func(T, T) bool) *TreeNode[T] {
	if len(slc) == 0 {
		return nil
	}

	slices.SortStableFunc(slc, func(a, b T) int {
		if equal(a, b) {
			return 0
		}
		if less(a, b) {
			return -1
		}
		return 1
	})

	var nodeInserter func([]T, int, int) *TreeNode[T]
	nodeInserter = func(t []T, start, end int) *TreeNode[T] {
		if start > end {
			return nil
		}

		mid := (start + end) / 2
		root := &TreeNode[T]{
			Val:   t[mid],
			Left:  nodeInserter(t, start, mid-1),
			Right: nodeInserter(t, mid+1, end),
		}

		return root
	}

	return nodeInserter(slc, 0, len(slc)-1)
}

func insertNode[T any](node *TreeNode[T], val T, less func(T, T) bool) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{Val: val}
	}

	if less(val, node.Val) {
		node.Left = insertNode(node.Left, val, less)
	} else {
		node.Right = insertNode(node.Right, val, less)
	}
	return node
}
