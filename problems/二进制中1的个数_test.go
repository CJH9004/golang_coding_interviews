package problems

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
