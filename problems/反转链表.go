// Description: 输入一个链表，反转链表后，输出新链表的表头。

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/list"
)

// ReverseList reverse List
func ReverseList(head *list.Node) *list.Node {
	if head == nil {
		return head
	}
	prev := head
	curr := head.Next
	if curr == nil {
		return head
	}
	next := curr.Next
	prev.Next = nil
	for next != nil {
		curr.Next = prev
		prev = curr
		curr = next
		next = next.Next
	}
	curr.Next = prev
	return curr
}
