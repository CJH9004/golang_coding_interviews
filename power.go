/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。
 * @solving:
 */
package main

import "math"

// Power receive a string and replace it's space by %20
func Power(base float64, n int) float64 {
	var ret float64 = 1
	if n == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	e := int(math.Abs(float64(n)))
	cur := base
	for e != 0 {
		if e&1 == 1 {
			ret *= cur
		}
		cur *= cur
		e >>= 1
	}
	return ret
}
