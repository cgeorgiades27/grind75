package addbinary

import (
	"testing"
)

func addBinaryCg(a, b string) string {
	getNumber := func(a string) int {
		num := 0
		for i := range a {
			num <<= 1
			if a[i] == '1' {
				num |= 1
			}
		}
		return num
	}

	au := getNumber(a)
	bu := getNumber(b)

	sum := au + bu
	var sumStr string
	for i := 0; i < 16; i++ {
		bit := "0"
		if (sum >> i & 1) == 1 {
			bit = "1"
		}
		sumStr = bit + sumStr
	}

	if len(sumStr) == 0 {
		return "0"
	}

	leading := len(sumStr) - 1
	for i := range sumStr {
		if sumStr[i] == '1' {
			leading = i
			break
		}
	}

	return sumStr[leading:]
}

func TestAddBinaryCg(t *testing.T) {
	for i, test := range TestCases {
		actual := addBinaryCg(test.a, test.b)
		if actual != test.expected {
			t.Errorf("Test #%d: got %s, want %s", i+1, actual, test.expected)
		}
	}
}
