// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。 输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
// 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。 NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

package problems

// MinNumberInRotateArray ...
func MinNumberInRotateArray(data []int) int {
	if len(data) == 0 {
		return 0
	}
	first := 0
	last := len(data) - 1
	for first < last {
		if data[first] < data[last] {
			return data[first]
		}
		mid := (last + first) >> 1
		if data[mid] > data[first] {
			first = mid + 1
		} else if data[mid] >= data[last] {
			last = mid
		} else {
			first++
		}
	}
	return data[first]
}
