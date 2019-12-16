// Description: 大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）。
// 使用闭包缓存已经计算的结果

package problems

// Fibonacci ...
var Fibonacci = func() func(n int) int {
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
