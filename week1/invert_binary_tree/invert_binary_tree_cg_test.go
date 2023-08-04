package invertbinarytree

import (
	"fmt"
	"testing"

	"github.com/cgeorgiades27/grind75/common"
)

func TestInvertTreeC(t *testing.T) {
	root := &common.TreeNode[int]{Val: 1}
	fmt.Println(InvertBinaryTreeC(root))
}

func TestInvertTreeCG(t *testing.T) {
	root := &common.TreeNode[int]{Val: 1}
	fmt.Println(InvertBinaryTreeCG(root))
}
