/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
 * @solving:
 */
package main

// RectCover receive a string and replace it's space by %20
var RectCover = func() func(n int) int {
	cache := []int{1, 1}
	return func(n int) int {
		if n < len(cache) {
			return cache[n]
		}
		for i := len(cache); i <= n; i++ {
			cache = append(cache, cache[i-1]+cache[i-2])
		}
		return cache[n]
	}
}()
