package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReConstructBinaryTree(t *testing.T) {
	type args struct {
		pre []int
		vin []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{pre: []int{1, 2, 4, 7, 3, 5, 6, 8}, vin: []int{4, 7, 2, 1, 5, 3, 8, 6}}, want: 1},
	}
	for _, tt := range tests {
		got := ReConstructBinaryTree(tt.args.pre, tt.args.vin)
		assert.Equal(t, tt.want, got.Val)
	}
}
