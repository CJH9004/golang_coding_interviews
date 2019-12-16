package problems

import (
	"testing"

	"github.com/CJH9004/golang_coding_interviews/kits/list"
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
		head := ReverseList(list.New(tt.data))
		ret := list.ToArray(head)
		assert.Equal(t, tt.ret, ret)
	}
}
