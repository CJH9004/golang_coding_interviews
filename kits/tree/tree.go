package tree

import (
	"github.com/CJH9004/golang_coding_interviews/kits/stack"
)

// Node for tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// NewNode return a Node
func NewNode(x int) *Node {
	return &Node{Val: x}
}

// PreOrder 前序遍历
func PreOrder(tree *Node, ret *[]int) {
	if tree == nil {
		return
	}
	*ret = append(*ret, tree.Val)
	PreOrder(tree.Left, ret)
	PreOrder(tree.Right, ret)
}

// SPreOrder 前序遍历 by stack
func SPreOrder(tree *Node, ret *[]int) {
	s := stack.New()
	s.Push(tree)
	for s.Len() > 0 {
		tmp, _ := s.Pop()
		node := tmp.(*Node)
		*ret = append(*ret, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
}

// InOrder 中序遍历
func InOrder(tree *Node, ret *[]int) {
	if tree == nil {
		return
	}
	InOrder(tree.Left, ret)
	*ret = append(*ret, tree.Val)
	InOrder(tree.Right, ret)
}

// PostOrder 后序遍历
func PostOrder(tree *Node, ret *[]int) {
	if tree == nil {
		return
	}
	PostOrder(tree.Left, ret)
	PostOrder(tree.Right, ret)
	*ret = append(*ret, tree.Val)
}
