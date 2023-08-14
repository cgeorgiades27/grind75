package validanagram

import "testing"

func TestValidAnagramCG(t *testing.T) {
	for i, test := range TestCases {
		actual := validAnagramCG(test.input1, test.input2)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}

func TestValidAnagramC(t *testing.T) {
	for i, test := range TestCases {
		actual := validAnagramC(test.input1, test.input2)
		if actual != test.output {
			t.Errorf("Test %d: expected %v, got %v", i+1, test.output, actual)
		}
	}
}

func BenchmarkValidAnagramCG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validAnagramCG("anagram", "nagaram")
	}
}

func BenchmarkValidAnagramC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validAnagramC("anagram", "nagaram")
	}
}
