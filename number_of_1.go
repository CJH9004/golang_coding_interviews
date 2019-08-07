/**
 * @Author: CJH9004
 * @Date: 2019-08-06 17:09:57
 * @LastEditors: CJH9004
 * @LastEditTime: 2019-08-06 17:23:51
 * @Description: 输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
 * @solving:
 */
package main

// NumberOf1 receive a string and replace it's space by %20
func NumberOf1(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		n = (n - 1) & n
	}
	return cnt
}

// NumberOf1Low receive a string and replace it's space by %20
func NumberOf1Low(n int) int {
	cnt := 0
	flag := 1
	for flag != 0 {
		if flag&n != 0 {
			cnt++
		}
		flag <<= 1
	}
	return cnt
}
