package invertbinarytree

import (
	"fmt"
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestInvertTreeC(t *testing.T) {
	root := &common.TreeNode[int]{Val: 1}
	fmt.Println(invertBinaryTreeCGC(root))
}

func TestInvertTreeCG(t *testing.T) {
	// root := &common.TreeNode[int]{Val: 1}

	for i, test := range TestCases {
		root := common.SliceToTree[int](test.input, func(i1, i2 int) bool { return i1 < i2 })
		res := invertBinaryTreeCG(root)

	}
}
