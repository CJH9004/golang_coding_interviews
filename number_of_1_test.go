/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test NumberOf1
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOf1(t *testing.T) {
	tests := []struct {
		data int
		ret  int
	}{
		{1, 1},
		{0, 0},
		{3, 2},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.ret, NumberOf1(tt.data))
		assert.Equal(t, tt.ret, NumberOf1Low(tt.data))
	}
}
