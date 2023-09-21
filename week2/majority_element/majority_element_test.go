package majorityelement

import "testing"

// O(n) runtime, O(1) space
func majorityElementCg(nums []int) int {
	count, leader := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == leader {
			count++
		} else {
			count--
		}
		if count == 0 {
			leader = nums[i]
			count = 1
		}
	}
	return leader
}

// LeetCode Results: Runtime - 17ms Beats 67.30% of users with Go, Memory - 6.68MB Beats 15.36% of users with Go
func TestMajorityElementCg(t *testing.T) {
	for i, test := range TestCases {
		actual := majorityElementCg(test.nums)
		if actual != test.expected {
			t.Errorf("Test #%d: got %d, want %d", i+1, actual, test.expected)
		}
	}
}

func BenchmarkMajorityElementCg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range TestCases {
			majorityElementCg(test.nums)
		}
	}
}
