// Description: 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
// solving: 当插入时，直接插入 stack1
// 当弹出时，当 stack2 不为空，弹出 stack2 栈顶元素，如果 stack2 为空，将 stack1 中的全部数逐个出栈入栈 stack2，再弹出 stack2 栈顶元素

package problems

import (
	"github.com/CJH9004/golang_coding_interviews/kits/stack"
)

// Queue is a FIFO datastruct
type Queue struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

// NewQueue return a queue
func NewQueue() *Queue {
	return &Queue{
		s1: stack.New(),
		s2: stack.New(),
	}
}

// Push push a val to queue
func (q *Queue) Push(x int) {
	q.s1.Push(x)
}

// Pop pop a val of queue
func (q *Queue) Pop() (ret int, ok bool) {
	if q.s2.Len() > 0 {
		ret, ok = q.s2.Pop()
	} else {
		for {
			v, ok := q.s1.Pop()
			if !ok {
				break
			}
			q.s2.Push(v)
		}
		ret, ok = q.s2.Pop()
	}

	return ret, ok
}
