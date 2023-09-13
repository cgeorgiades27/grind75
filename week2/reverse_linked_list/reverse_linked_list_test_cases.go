package reverselinkedlist

var TestCases = []struct {
	input    []int
	expected []int
}{
	{
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	},
	{
		input:    []int{1, 2},
		expected: []int{2, 1},
	},
	{
		input:    []int{},
		expected: []int{},
	},
}
