/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。 输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。 例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。 NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
 * @solving:
 */
package main

import "math"

// MinNumberInRotateArray receive a string and replace it's space by %20
func MinNumberInRotateArray(data []int) int {
	if len(data) == 0 {
		return 0
	}
	first := 0
	last := len(data) - 1
	mid := int(math.Ceil(float64(last+first) / 2))
	for ; first+1 < last; mid = int(math.Ceil(float64(last+first) / 2)) {
		if data[mid] < data[first] {
			last = mid
		}
		if data[mid] >= data[last] {
			first = mid
		}
	}
	return data[mid]
}
