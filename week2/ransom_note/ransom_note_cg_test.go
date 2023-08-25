package ransomnote

import "testing"

func canConstructCg(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	mColl := make(map[byte]int)
	for i := range magazine {
		mColl[magazine[i]]++
	}

	for i := range ransomNote {
		if mColl[ransomNote[i]] < 1 {
			return false
		}
		mColl[ransomNote[i]]--
	}

	return true
}

func TestCanConstructCg(t *testing.T) {
	for i, test := range testCases {
		actual := canConstructCg(test.ransomNote, test.magazine)
		if actual != test.expected {
			t.Errorf("Test %d: expected %t, actual %t", i, test.expected, actual)
		}
	}
}
