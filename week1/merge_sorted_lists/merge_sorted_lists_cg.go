package mergesortedlists

import "github.com/cgeorgiades27/grind75/common"

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
