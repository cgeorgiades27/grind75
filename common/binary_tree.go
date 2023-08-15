package common

import (
	"slices"
	"unsafe"
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

bool less(int a, int b)    { return a < b; };
bool greater(int a, int b) { return a > b; };

TreeNode *arr_to_tree(int *, size_t, comparitor);
TreeNode *insert_node(TreeNode *, int, comparitor);

TreeNode *arr_to_tree_left(int *arr, size_t size)  { return arr_to_tree(arr, size, less); }
TreeNode *arr_to_tree_right(int *arr, size_t size) { return arr_to_tree(arr, size, greater); }

bool check_trees(TreeNode *, TreeNode *);

TreeNode *arr_to_tree(int *arr, size_t size, comparitor cmp)
{
    if (!arr || size == 0)
        return NULL;

    TreeNode *root = NULL;
    for (int i = 0; i < size; ++i)
        root = insert_node(root, arr[i], cmp);

    return root;
}

TreeNode *insert_node(TreeNode *node, int val, comparitor cmp)
{
    if (!node)
    {
        node = malloc(sizeof(TreeNode));
        node->val = val;
        node->left = NULL;
        node->right = NULL;
        return node;
    }

    if (cmp(val, node->val))
        node->left = insert_node(node->left, val, cmp);
    else
        node->right = insert_node(node->right, val, cmp);

    return node;
}

bool check_trees(TreeNode *an, TreeNode *tn)
{
	if (!an && !tn)
		return true;

	if (!an || !tn)
		return false;

	if (an->val != tn->val)
		return false;

	return  (check_trees(an->left, tn->left) && check_trees(an->right, tn->right));
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
	}

	preorderRecurser(root)
	return slc
}

func SliceToTree[T any](slc []T, cmp func(T, T) bool) *TreeNode[T] {
	if len(slc) == 0 {
		return nil
	}

	var root *TreeNode[T] = nil
	for _, elem := range slc {
		root = insertNode(root, elem, cmp)
	}

	return root
}

func SliceToSearchTree[T any](slc []T, cmp func(T, T) bool, equal func(T, T) bool) *TreeNode[T] {
	if len(slc) == 0 {
		return nil
	}

	slices.SortStableFunc(slc, func(a, b T) int {
		if equal(a, b) {
			return 0
		}
		if cmp(a, b) {
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

func insertNode[T any](node *TreeNode[T], val T, cmp func(T, T) bool) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{Val: val}
	}

	if cmp(val, node.Val) {
		node.Left = insertNode(node.Left, val, cmp)
	} else {
		node.Right = insertNode(node.Right, val, cmp)
	}
	return node
}

func SliceToTreeC(slc []int, invert bool) *TreeNode[int] {
	arr := (*C.int)(unsafe.Pointer(&slc[0]))
	if invert {
		return (*TreeNode[int])(unsafe.Pointer(C.arr_to_tree_right(arr, C.size_t(len(slc)))))
	}

	return (*TreeNode[int])(unsafe.Pointer(C.arr_to_tree_left(arr, C.size_t(len(slc)))))
}

func CompareTreesC(an, tn *TreeNode[int]) bool {
	anc := (*C.TreeNode)(unsafe.Pointer(an))
	tnc := (*C.TreeNode)(unsafe.Pointer(tn))
	return (bool)(C.check_trees(anc, tnc))
}
