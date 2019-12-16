// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/list"
)

// Merge ...
func Merge(head1 *list.Node, head2 *list.Node) (nh *list.Node) {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	if head1.Val < head2.Val {
		nh = head1
		head1 = head1.Next
	} else {
		nh = head2
		head2 = head2.Next
	}

	cur := nh
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}

	for head1 != nil {
		cur.Next = head1
		head1 = head1.Next
		cur = cur.Next
	}
	for head2 != nil {
		cur.Next = head2
		head2 = head2.Next
		cur = cur.Next
	}
	return
}
