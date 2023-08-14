package twosum

import (
	"slices"
	"testing"
)

func twoSumCG(target int, nums []int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if res, ok := m[target-num]; ok {
			return []int{res, i}
		}
		m[num] = i
	}
	return []int{}
}

func TestTwoSumCG(t *testing.T) {
	for i, test := range TestCases {
		actual := twoSumCG(test.target, test.nums)
		// ensure consistent order
		slices.Sort(actual)
		slices.Sort(test.output)

		for j, num := range test.output {
			if num != actual[j] {
				t.Errorf("test %d failed, wanted: %d, got: %d", i, num, actual[j])
			}
		}
	}
}
