/**
 * @Author: CJH9004
 * @Date: 2019-08-07 14:35:29
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-07 14:35:37
 * @Description: test FindKthToTail
 */
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthToTail(t *testing.T) {
	tests := []struct {
		data []int
		k    int
		ret  int
		ok   bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4, true},
		{[]int{1, 2, 3, 4, 5}, 8, 0, false},
	}
	for _, tt := range tests {
		ret, ok := FindKthToTail(createList(tt.data), tt.k)
		if ok {
			assert.Equal(t, tt.ret, ret.val)
		}
		assert.Equal(t, tt.ok, ok)
	}
}
