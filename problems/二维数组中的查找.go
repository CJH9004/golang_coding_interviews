// Description: 在一个二维数组中（每个一维数组的长度相同），
// 每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，
// 判断数组中是否含有该整数。
// solving: 从二维数组左下角看，往右递增，往上递减

package problems

// FindIn2DArray receive a 2d array and a target, return if target in array
func FindIn2DArray(target int, array [][]int) bool {
	rows := len(array)
	if rows == 0 {
		return false
	}
	cols := len(array[0])

	for i, j := rows-1, 0; i >= 0 && j < cols; {
		if target > array[i][j] {
			j++
		} else if target < array[i][j] {
			i--
		} else {
			return true
		}
	}

	return false
}
