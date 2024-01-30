package binarysearch

/*
#include <stdlib.h>

int binary_search(int, int, int*);

int binary_search(int target, int arr_size, int *arr)
{
	int low = 0;
	int high = arr_size - 1;
	while (low <= high)
	{
		int mid = high - low / 2 + low;
		if (arr[mid] == target)
			return mid;
		if (target < arr[mid])
			high = mid - 1;
		else
			low = mid + 1;
	}
	return -1;
}
*/
import "C"
import "unsafe"

func binary_search_c(target int, slc []int) int {
	size := len(slc)
	arr := (*C.int)(unsafe.Pointer(&slc[0]))
	return (int)(C.binary_search(C.int(target), C.int(size), arr))
}

func binarySearchCg2(target int, slc []int) int {
	low, high := 0, len(slc)-1
	for low <= high {
		mid := high - low/2 + low
		if slc[mid] == target {
			return mid
		}
		if target < slc[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

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
