// list

package list

// Node is node for List
type Node struct {
	Val  int
	Next *Node
}

// New a List
func New(data []int) (head *Node) {
	if len(data) == 0 {
		return nil
	}
	head = &Node{data[0], nil}
	tail := head
	for _, v := range data[1:] {
		tail.Next = &Node{v, nil}
		tail = tail.Next
	}
	return
}

// ToArray transform list to array
func ToArray(head *Node) (ret []int) {
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return
}
