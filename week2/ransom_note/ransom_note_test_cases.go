package ransomnote

var testCases = []struct {
	ransomNote string
	magazine   string
	expected   bool
}{
	{
		ransomNote: "a",
		magazine:   "b",
		expected:   false,
	},
	{
		ransomNote: "aa",
		magazine:   "ab",
		expected:   false,
	},
	{
		ransomNote: "aa",
		magazine:   "aab",
		expected:   true,
	},
	{
		ransomNote: "bg",
		magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
		expected:   true,
	},
}
