package mergesortedlists

import (
	"fmt"
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestMergeSortedListsCG(t *testing.T) {
	for i, test := range TestCases {
		l1, l2 := common.SliceToList[int](test.input1), common.SliceToList[int](test.input2)
		actual := common.ListToSlice[int](mergeSortedListsCG(l1, l2))
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
