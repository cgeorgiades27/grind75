package linkedlistcycle

import (
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestHasCycleCG(t *testing.T) {
	for i, test := range TestCases {
		head := common.SliceToList[int](test.input)
		if test.index != -1 {
			hPtr, iPtr := head, head
			count := 0
			for hPtr.Next != nil {
				if count == test.index {
					iPtr = hPtr
				}
				count++
				hPtr = hPtr.Next
			}
			hPtr.Next = iPtr
		}
		actual := hasCycleCG(head)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, actual %v", i, test.output, actual)
		}
	}
}
