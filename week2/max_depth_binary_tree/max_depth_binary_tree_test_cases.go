package maxdepthbinarytree

var TestCases = []struct {
	input    []int
	expected int
}{
	{
		input:    []int{3, 9, 20, 0, 0, 15, 7},
		expected: 3,
	},
	{
		input:    []int{1, 0, 2},
		expected: 2,
	},
}
