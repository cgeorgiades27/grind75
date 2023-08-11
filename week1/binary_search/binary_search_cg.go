package binarysearch

func binarySearchCG(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
