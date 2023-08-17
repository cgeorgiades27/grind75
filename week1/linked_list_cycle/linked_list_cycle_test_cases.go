package linkedlistcycle

var TestCases = []struct {
	input  []int
	index  int
	output bool
}{
	{
		input:  []int{3, 2, 0, -4},
		index:  1,
		output: true,
	},
	{
		input:  []int{1, 2},
		index:  0,
		output: true,
	},
	{
		input:  []int{1},
		index:  -1,
		output: false,
	},
}
