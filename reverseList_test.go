/**
 * @Author: CJH9004
 * @Date: 2019-08-07 14:58:30
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-07 14:58:33
 * @Description: test ReverseList
 */
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		data []int
		ret  []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, nil},
	}
	for _, tt := range tests {
		head := ReverseList(createList(tt.data))
		ret := listToArray(head)
		assert.Equal(t, tt.ret, ret)
	}
}
