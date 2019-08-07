/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test MinNumberInRotateArray
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinNumberInRotateArray(t *testing.T) {
	tests := []struct {
		data []int
		ret  int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		ret := MinNumberInRotateArray(tt.data)
		assert.Equal(t, tt.ret, ret)
	}
}
