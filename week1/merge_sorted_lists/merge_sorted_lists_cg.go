package mergesortedlists

import "github.com/cgeorgiades27/grind75/common"

/*
#include <stdlib.h>

typedef struct ListNode {
	struct ListNode* next;
	int val;
} ListNode;

ListNode* merge_sorted_lists(ListNode* l1, ListNode* l2)
{
	ListNode* curr = NULL;
	ListNode* final = NULL;

	ListNode* itr1 = list1;
	ListNode* itr2 = list2;

	while (itr1 && itr2)
	{
		ListNode* temp = NULL;
		if (itr1->val < itr2->val)
		{
			temp = itr1;
			itr1 = itr1->next;
		}
		else
		{
			temp = itr2;
			itr2 = itr2->next;
		}

		// first look
		if (!final)
		{
			final = temp;
			curr = temp;
		}
		else
		{
			curr->next = temp;
			curr = curr->next;
		}
	}

	// drain remaining lists
	if (itr1)
	{
		if (!final)
			final->next = itr1;
		else
			curr->next = itr1;
	}

	if (itr2)
	{
		if (!final)
			final->next = itr2;
		else
			curr->next = itr2;
	}

	// head of merged list
	return final;
}

*/

func mergeSortedListsCG(list1, list2 *common.ListNode[int]) *common.ListNode[int] {
	var current, final *common.ListNode[int]
	itr1, itr2 := list1, list2

	for itr1 != nil && itr2 != nil {
		var temp *common.ListNode[int]
		if itr1.Val < itr2.Val {
			temp = itr1
			itr1 = itr1.Next
		} else {
			temp = itr2
			itr2 = itr2.Next
		}

		if final == nil {
			final = temp
			current = temp
		} else {
			current.Next = temp
			current = current.Next
		}

	}

	if itr1 != nil {
		if final == nil {
			final = itr1
		} else {
			current.Next = itr1
		}
	}

	if itr2 != nil {
		if final == nil {
			final = itr2
		} else {
			current.Next = itr2
		}
	}

	return final
}
