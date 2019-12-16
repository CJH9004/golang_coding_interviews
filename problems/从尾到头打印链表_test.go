package problems

import (
	"fmt"
	"testing"

	"github.com/CJH9004/golang_coding_interviews/kits/list"

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
		ret := ReversePrintList(list.New(tt.data))
		assert.Equal(
			t, tt.ret, ret, fmt.Sprintf(
				"Test ReversePrintList(%v); got %v; expect %v", tt.data, ret, tt.ret))
	}
}
