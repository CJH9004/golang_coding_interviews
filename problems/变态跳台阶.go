// Description: 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。

package problems

// JumpFloorII ...
var JumpFloorII = func() func(n int) int {
	cache := []int{1, 1}
	return func(n int) int {
		if n < len(cache) {
			return cache[n]
		}
		for i := len(cache); i <= n; i++ {
			cache = append(cache, 2*cache[i-1])
		}
		return cache[n]
	}
}()
