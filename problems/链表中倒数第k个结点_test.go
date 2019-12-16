package problems

import (
	"testing"

	"github.com/CJH9004/golang_coding_interviews/kits/list"
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
		ret, ok := FindKthToTail(list.New(tt.data), tt.k)
		if ok {
			assert.Equal(t, tt.ret, ret.Val)
		}
		assert.Equal(t, tt.ok, ok)
	}
}
