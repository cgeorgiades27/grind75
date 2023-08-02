package mergesortedlists

var TestCases = []struct {
	input1 []int
	input2 []int
	output []int
}{
	{
		input1: []int{1, 2, 4},
		input2: []int{1, 3, 4},
		output: []int{1, 1, 2, 3, 4, 4},
	},
	{
		input1: []int{},
		input2: []int{},
		output: []int{},
	},
	{
		input1: []int{},
		input2: []int{0},
		output: []int{0},
	},
}
