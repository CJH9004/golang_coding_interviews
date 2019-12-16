package tree

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
