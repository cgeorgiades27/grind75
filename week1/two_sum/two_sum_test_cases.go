package twosum

var TestCases = []struct {
	target int
	nums   []int
	output []int
}{
	{
		target: 9,
		nums:   []int{2, 7, 11, 15},
		output: []int{0, 1},
	},
	{
		target: 6,
		nums:   []int{3, 2, 4},
		output: []int{1, 2},
	},
	{
		target: 6,
		nums:   []int{3, 3},
		output: []int{0, 1},
	},
}
