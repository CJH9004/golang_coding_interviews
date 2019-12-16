package problems

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
