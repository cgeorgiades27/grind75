package diameterbinarytree

import (
	"testing"
)

func TestDiameterOfBinaryTreeCg(t *testing.T) {
	for i, test := range TestCases {
		actual := diameterOfBinaryTreeCg(test.input)
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}
