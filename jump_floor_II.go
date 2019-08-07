/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
 * @solving:
 */
package main

// JumpFloorII receive a string and replace it's space by %20
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