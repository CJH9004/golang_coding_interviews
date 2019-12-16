package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpFloorII(t *testing.T) {
	tests := []struct {
		data int
		ret  int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 8},
		{5, 16},
		{6, 32},
	}

	for _, tt := range tests {
		ret := JumpFloorII(tt.data)
		assert.Equal(t, tt.ret, ret)
	}
}
