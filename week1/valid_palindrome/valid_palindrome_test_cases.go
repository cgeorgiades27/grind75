package validpalindrome

var TestCases = []struct {
	input  string
	output bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
	{".,", true},
}
