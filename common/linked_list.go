package common

/*
#include <stdlib.h>

typedef struct ListNode
{
	struct ListNode* next;
	int val;
} ListNode;


ListNode* array_to_list(int *arr, int size)
{
	if (size == 0)
		return NULL;

	ListNode* head = malloc(sizeof(ListNode));
	head->val = arr[0];
	ListNode* curr = head;

	for (int i = 1; i < size; ++i)
	{
		curr->next = malloc(sizeof(ListNode));
		curr = curr->next;
		curr->val = arr[i];
	}

	curr->next = NULL;

	return head;
}
*/
import "C"
import "unsafe"

func SliceToListC(slc []int) *ListNode[int] {
	size := len(slc)
	arr := (*C.int)(unsafe.Pointer(&slc[0]))

	head := C.array_to_list(arr, C.int(size))
	return (*ListNode[int])(unsafe.Pointer(head))
}

// Definition for singly-linked list.
type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func ListToSlice[T any](head *ListNode[T]) []T {
	var slc []T
	for head != nil {
		slc = append(slc, head.Val)
		head = head.Next
	}
	return slc
}

func SliceToList[T any](slc []T) *ListNode[T] {
	if len(slc) == 0 {
		return nil
	}

	head := &ListNode[T]{}
	curr := head
	for i := 0; i < len(slc); i++ {
		curr.Val = slc[i]
		if i < len(slc)-1 {
			curr.Next = &ListNode[T]{}
			curr = curr.Next
		}
	}
	return head
}
