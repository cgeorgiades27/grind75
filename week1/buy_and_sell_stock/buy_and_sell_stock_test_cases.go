package buyandsellstock

var TestCases = []struct {
	input  []int
	output int
}{
	{
		input:  []int{7, 1, 5, 3, 6, 4},
		output: 5,
	},
	{
		input:  []int{7, 6, 4, 3, 1},
		output: 0,
	},
}
