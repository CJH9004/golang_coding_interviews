// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
// 例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
// 则依次打印出数字1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10.

package problems

// PrintMatrix ...
func PrintMatrix(m [][]int) []int {
	if len(m) == 0 || len(m[0]) == 0 {
		return []int{}
	}

	r, c := len(m), len(m[0])
	result := make([]int, 0, r*c)
	for i := 0; i < c-1; i++ {
		result = append(result, m[0][i])
	}
	for i := 0; i < r-1; i++ {
		result = append(result, m[i][c-1])
	}
	for i := c - 1; i > 0; i-- {
		result = append(result, m[r-1][i])
	}
	for i := r - 1; i > 0; i-- {
		result = append(result, m[i][0])
	}
	tmp := make([][]int, r-2)
	for i := 1; i < r-1; i++ {
		tmp[i-1] = make([]int, c-2)
		for j := 1; j < c-1; j++ {
			tmp[i-1][j-1] = m[i][j]
		}
	}
	return append(result, PrintMatrix(tmp)...)
}
