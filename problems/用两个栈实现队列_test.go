/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test Queue
 * TODO:
 */

package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	_, ok := q.Pop()
	assert.False(t, ok)

	q.Push(1)
	q.Push(2)
	v, ok := q.Pop()
	q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Pop()
	q.Pop()
	q.Push(1)
	v, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
}
