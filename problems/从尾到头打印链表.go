// 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/list"
)

// ReversePrintList ...
func ReversePrintList(head *list.Node) (data []int) {
	// for cur := head; cur != nil; cur = cur.next {
	// 	data = append([]int{cur.val}, data...)
	// }

	// 获取长度
	cnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}
	// 反向给数组赋值
	data = make([]int, cnt)
	for cur := head; cur != nil; cur = cur.Next {
		cnt--
		data[cnt] = cur.Val
	}
	return
}
