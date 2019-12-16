// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。

package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CJH9004/golang_coding_interviews/kits/list"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		arg1 []int
		arg2 []int
		want []int
	}{
		{[]int{1, 3, 5, 7, 10}, []int{1, 2, 8, 9, 12, 14, 15, 18}, []int{1, 1, 2, 3, 5, 7, 8, 9, 10, 12, 14, 15, 18}},
	}
	for _, tt := range tests {
		act := list.ToArray(Merge(list.New(tt.arg1), list.New(tt.arg2)))
		assert.Equal(t, tt.want, act)
	}
}
