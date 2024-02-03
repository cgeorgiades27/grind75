package validpalindrome

import (
	"strings"
)

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for !isAlpha(rune(s[left])) &&
			left < right {
			left++
		}
		for !isAlpha(rune(s[right])) &&
			right > left {
			right--
		}

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isIntPalindrome(num int) bool {
	var nums []int
	i := 1
	for num%i < num {
		i *= 10
		nums = append(nums, num%i/(i/10))
	}

	first, last := 0, len(nums)-1
	for first < last {
		if nums[first] != nums[last] {
			return false
		}
		first++
		last--
	}

	return true
}
