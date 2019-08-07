/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。
 * @solving:
 */
package main

// ReversePrintList receive a string and replace it's space by %20
func ReversePrintList(head *Node) (data []int) {
	// for cur := head; cur != nil; cur = cur.next {
	// 	data = append([]int{cur.val}, data...)
	// }
	cnt := 0
	for cur := head; cur != nil; cur = cur.next {
		cnt++
	}
	data = make([]int, cnt)
	for cur := head; cur != nil; cur = cur.next {
		cnt--
		data[cnt] = cur.val
	}
	return
}
