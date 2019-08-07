/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
 * @solving: 循环队列
 */
package main

// Queue is a FILO data struct
type Queue interface {
	Push(n int) bool
	Pop() (int, bool)
}

type queue struct {
	data []int
	head int
	tail int
	len  int
}

// NewQueue return a Queue
func NewQueue(d []int) Queue {
	dLen := len(d)
	q := queue{}
	q.len = dLen
	if dLen == 0 {
		dLen = 1
		q.head = -1
		q.tail = -1
	} else {
		q.head = dLen
		q.tail = dLen*2 - 1
	}
	q.data = make([]int, dLen*2)
	copy(q.data[dLen:], d)

	return &q
}

func (q *queue) Push(n int) bool {
	i := q.head - 1
	if i < 0 {
		i = len(q.data) - 1
	}
	if q.tail == i {
		l := len(q.data)
		d := make([]int, l*2)
		copy(d[len(q.data):], q.data)
		q.data = d
		i = l - 1
		q.tail = len(q.data) - 1
	}
	q.data[i] = n
	q.head = i
	if q.tail < 0 {
		q.tail = i
	}
	q.len++
	return true
}

func (q *queue) Pop() (int, bool) {
	if q.len <= 0 || q.tail < 0 {
		return 0, false
	}
	d := q.data[q.tail]
	q.tail--
	if q.tail < 0 && q.len > 0 {
		q.tail = len(q.data) - 1
	}
	q.len--
	if q.len <= 0 {
		q.tail = -1
		q.head = -1
	}
	return d, true
}
