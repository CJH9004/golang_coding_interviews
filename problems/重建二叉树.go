// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
// 前序遍历：root - 前序遍历左子树 - 前序遍历右子树
// 中序遍历：中序遍历左子树 - root - 中序遍历右子树
// 后序遍历：后序遍历左子树 - 后序遍历右子树 - root
// 前中后值得是root位置，左子树优先于右子树

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/tree"
)

// ReConstructBinaryTree ...
func ReConstructBinaryTree(pre, vin []int) *tree.Node {
	if len(pre) == 0 || len(pre) != len(vin) {
		return nil
	}

	root := tree.NewNode(pre[0])
	i := indexOf(vin, pre[0])
	if i == -1 {
		return nil
	}

	root.Left = ReConstructBinaryTree(pre[1:i+1], vin[0:i])
	root.Right = ReConstructBinaryTree(pre[i+1:], vin[i+1:])
	return root
}

func indexOf(a []int, t int) int {
	for i := range a {
		if a[i] == t {
			return i
		}
	}
	return -1
}
