package threesum

var TestCases = []struct {
	Input    []int
	Expected [][]int
}{
	{
		Input:    []int{-1, 0, 1, 2, -1, -4},
		Expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	},
	{
		Input:    []int{0, 1, 1},
		Expected: [][]int{},
	},
	{
		Input:    []int{0, 0, 0},
		Expected: [][]int{{0, 0, 0}},
	},
}
