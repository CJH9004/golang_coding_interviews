package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReOrderArray(t *testing.T) {
	tests := []struct {
		data []int
		ret  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		ReOrderArray(&tt.data)
		assert.Equal(t, tt.ret, tt.data)
	}
}
