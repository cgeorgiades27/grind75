package addbinary

import (
	"testing"
)

func addBinaryCg(a, b string) string {
	var au, bu int
	for i := range a {
		if a[i] == '1' {
			au = au | 1
		}
		if i < len(a)-1 {
			au = au << 1
		}
	}

	for i := range b {
		if b[i] == '1' {
			bu = bu | 1
		}

		if i < len(b)-1 {
			bu = bu << 1
		}
	}

	sum := au + bu
	var sumStr string
	for i := 0; i < 16; i++ {
		bit := "0"
		if (sum >> i & 1) == 1 {
			bit = "1"
		}
		sumStr = bit + sumStr
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
