// 请实现一个函数，将一个字符串中的每个空格替换成“%20”。例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
// 遍历string

package problems

// ReplaceSpace receive a string and replace it's space by %20
func ReplaceSpace(str string) string {
	ret := ""
	for _, v := range str {
		if v == ' ' {
			ret += "%20"
		} else {
			ret += string(v)
		}
	}
	return ret
}
