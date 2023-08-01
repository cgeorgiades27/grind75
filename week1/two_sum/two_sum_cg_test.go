package week1

import (
	"slices"
	"testing"
)

func TwoSumCG(target int, nums []int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if res, ok := m[target-num]; ok {
			return []int{res, i}
		}
		m[target-num] = i
	}
	return []int{}
}

func TestTwoSumCG(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		output []int
	}{
		{target: 9, nums: []int{2, 7, 11, 15}, output: []int{0, 1}},
		{target: 6, nums: []int{3, 2, 4}, output: []int{1, 2}},
		{target: 6, nums: []int{3, 3}, output: []int{0, 1}},
	}

	for i, test := range tests {
		actual := TwoSumCG(test.target, test.nums)
		// ensure consistent order
		slices.Sort(actual)
		slices.Sort(test.output)

		for j, num := range actual {
			if num != test.output[j] {
				t.Errorf("test %d failed, wanted: %d, got: %d", i, test.output[j], num)
			}
		}
	}
}
