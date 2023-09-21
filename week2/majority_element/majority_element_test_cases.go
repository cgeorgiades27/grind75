package majorityelement

var TestCases = []struct {
	nums     []int
	expected int
}{
	{
		nums:     []int{3, 2, 3},
		expected: 3,
	},
	{
		nums:     []int{2, 2, 1, 1, 1, 2, 2},
		expected: 2,
	},
}