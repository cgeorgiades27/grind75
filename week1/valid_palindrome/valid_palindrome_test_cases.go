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
