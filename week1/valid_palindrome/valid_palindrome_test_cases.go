package validpalindrome

var TestCases = []struct {
	input  string
	output bool
}{
	{
		input:  "A man, a plan, a canal: Panama",
		output: true,
	},
	{
		input:  "race a car",
		output: false,
	},
	{
		input:  " ",
		output: true,
	},
	{
		input:  ".,",
		output: true,
	},
}

var TestCasesInt = []struct {
	input    int
	expected bool
}{
	{
		input:    121,
		expected: true,
	},
	{
		input:    1234,
		expected: false,
	},
	{
		input:    12321,
		expected: true,
	},
}
