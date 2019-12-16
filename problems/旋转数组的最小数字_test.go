package problems

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
