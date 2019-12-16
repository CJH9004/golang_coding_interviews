// Description: 输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。

package problems

// NumberOf1 n有多少1，就循环多少次
func NumberOf1(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		n = (n - 1) & n
	}
	return cnt
}

// NumberOf1Low n有多少位，就循环多少次
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
