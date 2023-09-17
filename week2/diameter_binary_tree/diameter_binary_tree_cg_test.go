package diameterbinarytree

import (
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestDiameterOfBinaryTreeC(t *testing.T) {
	for i, test := range TestCases {
		actual := diameterOfBinaryTreeC(common.SliceToTreeC(test.input, false))
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}

func TestDiameterOfBinaryTreeCg(t *testing.T) {
	for i, test := range TestCases {
		actual := diameterOfBinaryTreeCg(common.SliceToTree(test.input, func(i1, i2 int) bool { return i1 < i2 }))
		if actual != test.expected {
			t.Errorf("Test %d: expected %d, got %d", i+1, test.expected, actual)
		}
	}
}
