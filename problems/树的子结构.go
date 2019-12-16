// 输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/tree"
)

// HasSubtree ...
func HasSubtree(root1 *tree.Node, root2 *tree.Node) (result bool) {
	if root1 != nil && root2 != nil { //当Tree1和Tree2都不为零的时候，才进行比较。否则直接返回false
		if root1.Val == root2.Val { //如果找到了对应Tree2的根节点的点
			//以这个根节点为为起点判断是否包含Tree2
			result = doesTree1HaveTree2(root1, root2)
		}
		if !result { //如果找不到，那么就再去root的左儿子当作起点，去判断时候包含Tree2
			result = HasSubtree(root1.Left, root2)
		}
		if !result { //如果还找不到，那么就再去root的右儿子当作起点，去判断时候包含Tree2
			result = HasSubtree(root1.Right, root2)
		}
	}
	return
}

func doesTree1HaveTree2(root1 *tree.Node, root2 *tree.Node) (result bool) {
	if root2 == nil { //如果Tree2已经遍历完了都能对应的上，返回true
		return true
	}
	if root1 == nil { //如果Tree2还没有遍历完，Tree1却遍历完了。返回false
		return false
	}
	if root1.Val != root2.Val { //如果其中有一个点没有对应上，返回false
		return false
	}
	return doesTree1HaveTree2(root1.Left, root2.Left) && doesTree1HaveTree2(root1.Right, root2.Right)
}
