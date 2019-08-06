/**
 * @Author: CJH9004
 * @Date: 2019-08-06 16:06:50
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:10:28
 * @Description: test reversePrintList
 */

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrintList(t *testing.T) {
	tests := []struct {
		data []int
		ret  []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		ret := ReversePrintList(createList(tt.data))
		assert.Equal(
			t, tt.ret, ret, fmt.Sprintf(
				"Test ReversePrintList(%v); got %v; expect %v", tt.data, ret, tt.ret))
	}
}
