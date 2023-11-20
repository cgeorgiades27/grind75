package binarysearch

// binarySearchCg is a binary search implementation that takes a target int and a sorted int slice.
func binarySearchCg(target int, slc []int) int {
	low, high := 0, len(slc)-1

	for low <= high {
		mid := (high-low)/2 + low
		if slc[mid] == target {
			return mid
		}
		if target > slc[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
