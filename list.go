/**
 * @Author: CJH9004
 * @Date: 2019-08-07 14:30:14
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-07 14:30:26
 * @Description: List
 */
package main

// Node is node for List
type Node struct {
	val  int
	next *Node
}

func createList(data []int) (head *Node) {
	if len(data) == 0 {
		return nil
	}
	head = &Node{data[0], nil}
	tail := head
	for _, v := range data[1:] {
		tail.next = &Node{v, nil}
		tail = tail.next
	}
	return
}

func listToArray(head *Node) (ret []int) {
	for head != nil {
		ret = append(ret, head.val)
		head = head.next
	}
	return
}
