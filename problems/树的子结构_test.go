// 输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

package problems

import (
	"testing"

	"github.com/CJH9004/golang_coding_interviews/kits/tree"
	"github.com/stretchr/testify/assert"
)

func TestHasSubtree(t *testing.T) {
	root := tree.NewNode(0)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(2)
	root.Right.Left = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	root.Right.Right.Left = tree.NewNode(5)

	tests := []struct {
		arg1 *tree.Node
		arg2 *tree.Node
		want bool
	}{
		{root, root.Right, true},
		{root.Right, root.Left, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, HasSubtree(tt.arg1, tt.arg2))
	}
}
