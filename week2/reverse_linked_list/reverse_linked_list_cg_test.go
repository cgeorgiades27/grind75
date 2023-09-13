package reverselinkedlist

import (
	"slices"
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

// reverseLinkedListCg reverses a linked list in place iteratively.
func reverseLinkedListCg(head *common.ListNode[int]) *common.ListNode[int] {
	var curr, prev *common.ListNode[int]
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// reverseLinkedListRecCg reverses a linked list in place recursively.
func reverseLinkedListRecCg(head *common.ListNode[int]) *common.ListNode[int] {
	var recurse func(curr, prev *common.ListNode[int]) *common.ListNode[int]
	recurse = func(curr, prev *common.ListNode[int]) *common.ListNode[int] {
		if curr == nil {
			return prev
		}

		next := curr.Next
		curr.Next = prev
		return recurse(next, curr)
	}

	var prev *common.ListNode[int]
	return recurse(head, prev)
}

// LeetCode Results: Runtime 1 ms Beats 74.54% | Memory 2.6 MB Beats 59.12%
func TestReverseLinkedListCg(t *testing.T) {
	for i, test := range TestCases {
		actual := reverseLinkedListCg(common.SliceToList[int](test.input))
		actualInts := common.ListToSlice[int](actual)
		if !slices.Equal(test.expected, actualInts) {
			t.Errorf("Test %d: expected %v, actual %v", i+1, test.expected, actualInts)
		}
	}
}

// LeetCode Results: Runtime 3 ms Beats 27.72% | Memory 3 MB Beats 8.6%
func TestReverseLinkedListRecCg(t *testing.T) {
	for i, test := range TestCases {
		actual := reverseLinkedListRecCg(common.SliceToList[int](test.input))
		actualInts := common.ListToSlice[int](actual)
		if !slices.Equal(test.expected, actualInts) {
			t.Errorf("Test %d: expected %v, actual %v", i+1, test.expected, actualInts)
		}
	}
}
