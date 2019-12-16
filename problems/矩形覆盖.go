// Description: 我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？

package problems

// RectCover ...
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
