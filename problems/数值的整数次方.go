// Description: 给定一个double类型的浮点数base和int类型的整数exponent。求base的exponent次方。

package problems

// Power ...
func Power(base float64, n int) float64 {
	var ret float64 = 1
	if n == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	e := n
	if e < 0 {
		e = -n
	}
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
