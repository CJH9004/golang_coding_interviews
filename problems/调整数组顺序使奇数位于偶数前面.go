// Description: 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

package problems

// ReOrderArray reorder a array
func ReOrderArray(a *[]int) {
	t := make([]int, len(*a))
	i := 0
	for _, v := range *a {
		if v%2 == 1 {
			t[i] = v
			i++
		}
	}
	for _, v := range *a {
		if v%2 == 0 {
			t[i] = v
			i++
		}
	}
	*a = t
}
