package main

import "fmt"

/*
*
给你一个字符串 s 和一个整数 t，表示要执行的 转换 次数。每次 转换 需要根据以下规则替换字符串 s 中的每个字符：

如果字符是 'z'，则将其替换为字符串 "ab"。
否则，将其替换为字母表中的下一个字符。例如，'a' 替换为 'b'，'b' 替换为 'c'，依此类推。
返回 恰好 执行 t 次转换后得到的字符串的 长度。

由于答案可能非常大，返回其对 109 + 7 取余的结果。

示例 1：

输入： s = "abcyy", t = 2

输出： 7

解释：

第一次转换 (t = 1)
'a' 变为 'b'
'b' 变为 'c'
'c' 变为 'd'
'y' 变为 'z'
'y' 变为 'z'
第一次转换后的字符串为："bcdzz"
第二次转换 (t = 2)
'b' 变为 'c'
'c' 变为 'd'
'd' 变为 'e'
'z' 变为 "ab"
'z' 变为 "ab"
第二次转换后的字符串为："cdeabab"
最终字符串长度：字符串为 "cdeabab"，长度为 7 个字符。

100

a

25  ab   75
*/
func lengthAfterTransformations(s string, t int) int {

	mod := int(1e9 + 7)

	ans := 0

	f := make([]map[int]int, 26)

	for i := range f {
		f[i] = make(map[int]int)
	}

	var cul func(a int, t int) int

	cul = func(a int, t int) int {
		if _, ok := f[a][t]; ok {
			return f[a][t]
		}
		cur := 0
		if a+t > 25 {
			cur = (cur + cul(0, t-(26-a)) + cul(1, t-(26-a)) + 1) % mod
		}
		f[a][t] = cur
		return cur
	}

	for _, v := range s {
		ans = (ans + cul(int(v-'a'), t)) % mod
	}

	return (ans + len(s)) % mod
}

func main() {
	fmt.Println(lengthAfterTransformations("abcyy", 100))
}
