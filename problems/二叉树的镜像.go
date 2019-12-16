// 操作给定的二叉树，将其变换为源二叉树的镜像。

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/tree"
)

// Mirror ...
func Mirror(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	Mirror(root.Left)
	Mirror(root.Right)
	return root
}
