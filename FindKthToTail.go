/**
 * @Author: CJH9004
 * @Date: 2019-08-07 14:27:53
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-07 14:27:56
 * @Description:
 */
package main

// FindKthToTail find kth element to tail in List
func FindKthToTail(head *Node, k int) (*Node, bool) {
	afterK := head
	for i := 0; i < k; i++ {
		if afterK == nil {
			return nil, false
		}
		afterK = afterK.next
	}
	for afterK != nil {
		head = head.next
		afterK = afterK.next
	}
	return head, true
}
