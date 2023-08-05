package binarysearch

var TestCases = []struct {
	nums   []int
	target int
	output int
}{
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 9,
		output: 4,
	},
	{
		nums:   []int{-1, 0, 3, 5, 9, 12},
		target: 2,
		output: -1,
	},
}
