package validparentheses

var TestCases = []struct {
	input  string
	output bool
}{
	{
		input:  "()",
		output: true,
	},
	{
		input:  "()[]{}",
		output: true,
	},
	{
		input:  "(]",
		output: false,
	},
	{
		input:  "()[][[[",
		output: false,
	},
}
