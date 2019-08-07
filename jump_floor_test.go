/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test JumpFloor
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpFloor(t *testing.T) {
	tests := []struct {
		data int
		ret  int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
	}

	for _, tt := range tests {
		ret := JumpFloor(tt.data)
		assert.Equal(t, tt.ret, ret)
	}
}
