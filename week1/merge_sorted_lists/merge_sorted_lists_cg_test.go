package mergesortedlists

import (
	"fmt"
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func MergeSortedListsCG(list1, list2 *common.ListNode[int]) *common.ListNode[int] {
	var current, final *common.ListNode[int]
	itr1, itr2 := list1, list2

	final = nil
	current = nil

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

func TestMergeSortedListsCG(t *testing.T) {
	for i, test := range TestCases {
		l1, l2 := common.SliceToList[int](test.input1), common.SliceToList[int](test.input2)
		actual := common.ListToSlice[int](MergeSortedListsCG(l1, l2))
		fmt.Println(actual, test.output)
		if len(test.output) != len(actual) {
			t.Fatalf("test %d, wanted len of %d, got len of %d", i, len(test.output), len(actual))
		}

		for j, num := range test.output {
			if num != actual[j] {
				t.Fatalf("test %d, wanted %d at index %d, got %d", i, num, j, actual[j])
			}
		}
	}
}
