// 输入一个链表，输出该链表中倒数第k个结点。

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/list"
)

// FindKthToTail find kth element to tail in List
func FindKthToTail(head *list.Node, k int) (*list.Node, bool) {
	afterK := head
	for i := 0; i < k; i++ {
		if afterK == nil {
			return nil, false
		}
		afterK = afterK.Next
	}
	for afterK != nil {
		head = head.Next
		afterK = afterK.Next
	}
	return head, true
}
