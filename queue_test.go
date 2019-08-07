/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test Queue
 * TODO:
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		init     []int
		pushVal  int
		pushRet  bool
		popRet   int
		popRetOk bool
	}{
		{[]int{1, 2, 3, 4}, 5, true, 4, true},
		{[]int{1}, 5, true, 1, true},
		{[]int{}, 5, true, 5, true},
	}

	for _, tt := range tests {
		q := NewQueue(tt.init)
		pushRet := q.Push(tt.pushVal)
		popRet, popRetOk := q.Pop()
		assert.Equal(t, tt.pushRet, pushRet)
		assert.Equal(t, tt.popRet, popRet)
		assert.Equal(t, tt.popRetOk, popRetOk)
	}
	q := NewQueue([]int{})
	_, popRetOk := q.Pop()
	assert.Equal(t, false, popRetOk)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	popRet, popRetOk := q.Pop()
	assert.Equal(t, 1, popRet)
	assert.Equal(t, true, popRetOk)
	q.Push(4)
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	_, popRetOk = q.Pop()
	assert.Equal(t, false, popRetOk)
}
