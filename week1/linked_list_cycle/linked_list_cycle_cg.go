package linkedlistcycle

/*
#include <stdio.h>
#include <stdbool.h>

 typedef struct ListNode
 {
	struct ListNode* next;
	int val;
 } ListNode;

bool has_cycle(ListNode* head)
{
   if (!head)
		return false;

	ListNode* slow = head;
	ListNode* fast = head;

	while (fast)
	{
		if (fast->next == slow)
			return true;

        slow = slow->next;
		if (!fast->next)
			fast = fast->next;
		else
			fast = fast->next->next;
	}
	return false;
}
*/
import "C"
import "github.com/cgeorgiades27/grind75/common"

func hasCycleCG(head *common.ListNode[int]) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil {
		if fast.Next == slow {
			return true
		}
		slow = slow.Next
		if fast.Next == nil {
			fast = fast.Next
		} else {
			fast = fast.Next.Next
		}
	}
	return false
}
