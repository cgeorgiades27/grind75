package validanagram

var TestCases = []struct {
	input1 string
	input2 string
	output bool
}{
	{
		input1: "anagram",
		input2: "nagaram",
		output: true,
	},
	{
		input1: "rat",
		input2: "car",
		output: false,
	},
}
