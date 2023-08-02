package common

import "testing"

func TestListToSlice(t *testing.T) {
	var list1 *ListNode[int] = nil
	var list2 *ListNode[int] = nil
	var list3 *ListNode[int] = nil

	list1 = &ListNode[int]{
		Val: 1,
		Next: &ListNode[int]{
			Val: 2,
			Next: &ListNode[int]{
				Val: 3,
				Next: &ListNode[int]{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	list2 = &ListNode[int]{Val: 1, Next: nil}

	tests := []struct {
		input  *ListNode[int]
		output []int
	}{
		{
			input:  list1,
			output: []int{1, 2, 3, 4},
		},
		{
			input:  list2,
			output: []int{1},
		},
		{
			input:  list3,
			output: []int{},
		},
	}

	for i, test := range tests {
		actual := ListToSlice[int](test.input)
		if len(actual) != len(test.output) {
			t.Errorf("test %d failed, wanted len of: %d, got: %d", i, len(test.output), len(actual))
		}

		for j, elem := range test.output {
			if elem != actual[j] {
				t.Errorf("test %d failed, wanted: %d, got: %d", i, elem, actual[j])
			}
		}
	}
}

func TestSliceToList(t *testing.T) {
	var list1 *ListNode[int] = nil
	var list2 *ListNode[int] = nil
	var list3 *ListNode[int] = nil

	list1 = &ListNode[int]{
		Val: 1,
		Next: &ListNode[int]{
			Val: 2,
			Next: &ListNode[int]{
				Val: 3,
				Next: &ListNode[int]{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	list2 = &ListNode[int]{Val: 1, Next: nil}

	tests := []struct {
		input  []int
		output *ListNode[int]
	}{
		{
			input:  []int{1, 2, 3, 4},
			output: list1,
		},
		{
			input:  []int{1},
			output: list2,
		},
		{
			input:  []int{},
			output: list3,
		},
	}

	for i, test := range tests {
		actual := SliceToList[int](test.input)
		for actual != nil && test.output != nil {
			if actual.Val != test.output.Val {
				t.Errorf("test %d failed, wanted val of: %d, got: %d", i, test.output.Val, actual.Val)
			}
			actual = actual.Next
			test.output = test.output.Next
		}
		if actual != nil {
			t.Errorf("test %d failed, actual is longer than test case", i)
		}
		if test.output != nil {
			t.Errorf("test %d failed, test case is longer than actual", i)
		}
	}
}
