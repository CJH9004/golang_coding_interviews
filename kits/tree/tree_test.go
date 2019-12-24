package tree

import (
	"fmt"
	"sync"
	"testing"
)

var (
	tree *Node
	once sync.Once
)

func treeSetup() {
	tree = NewNode(1)
	{
		tree.Left = NewNode(2)
		{
			tree.Left.Left = NewNode(4)
			tree.Left.Right = NewNode(5)
			{
				tree.Left.Right.Left = NewNode(6)
			}
		}
		tree.Right = NewNode(3)
		{
			tree.Right.Left = NewNode(4)
			tree.Right.Right = NewNode(5)
			{
				tree.Right.Left.Left = NewNode(7)
				tree.Right.Left.Right = NewNode(8)
			}
		}
	}
}

func TestPreOrder(t *testing.T) {
	once.Do(treeSetup)

	ret := make([]int, 0)
	PreOrder(tree, &ret)
	fmt.Println(ret)
}

func TestSPreOrder(t *testing.T) {
	once.Do(treeSetup)

	ret := make([]int, 0)
	SPreOrder(tree, &ret)
	fmt.Println(ret)
}

func TestInOrder(t *testing.T) {
	once.Do(treeSetup)

	ret := make([]int, 0)
	InOrder(tree, &ret)
	fmt.Println(ret)
}

func TestPostOrder(t *testing.T) {
	once.Do(treeSetup)

	ret := make([]int, 0)
	PostOrder(tree, &ret)
	fmt.Println(ret)
}
