package validanagram

import "testing"

func TestValidAnagramCG(t *testing.T) {
	for i, test := range TestCases {
		actual := ValidAnagramCG(test.input1, test.input2)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}

func TestValidAnagramC(t *testing.T) {
	for i, test := range TestCases {
		actual := ValidAnagramC(test.input1, test.input2)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}
