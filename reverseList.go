/**
 * @Author: CJH9004
 * @Date: 2019-08-07 14:51:22
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-07 14:51:25
 * @Description: 输入一个链表，反转链表后，输出新链表的表头。
 */
package main

// ReverseList reverse List
func ReverseList(head *Node) *Node {
	if head == nil {
		return head
	}
	prev := head
	curr := head.next
	if curr == nil {
		return head
	}
	next := curr.next
	prev.next = nil
	for next != nil {
		curr.next = prev
		prev = curr
		curr = next
		next = next.next
	}
	curr.next = prev
	return curr
}
