package threesum

import (
	"slices"

	"github.com/cgeorgiades27/grind75/common"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int
	low, high, sum := 0, 0, 0

	for i := range nums {
		low = i + 1
		high = len(nums) - 1
		for low < high {
			sum = nums[low] + nums[i] + nums[high]
			if sum == 0 {
				result = append(result, []int{nums[low], nums[i], nums[high]})
			}
			if sum < 0 {
				low++
			} else {
				high--
			}
		}
	}

	for i := 0; i < len(result)-1; i++ {
		if common.CheckSliceEquality[int](result[i], result[i+1]) {
			result = append(result[:i], result[i+1:]...)
		}
	}
	return result
}
