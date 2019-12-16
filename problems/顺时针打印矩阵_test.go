// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，

// 例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16

// 则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.

package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintMatrix(t *testing.T) {
	tests := []struct {
		arg  [][]int
		want []int
	}{
		{
			arg: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, PrintMatrix(tt.arg))
	}
}
