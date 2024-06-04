package main

import (
	"fmt"
	"strings"
)

/*
*给你一个二进制串 s  （一个只包含 0 和 1 的字符串），我们可以将 s 分割成 3 个 非空 字符串 s1, s2, s3 （s1 + s2 + s3 = s）。

请你返回分割 s 的方案数，满足 s1，s2 和 s3 中字符 '1' 的数目相同。

由于答案可能很大，请将它对 10^9 + 7 取余后返回。

输入：s = "0000"
输出：3
解释：总共有 3 种分割 s 的方法。
"0|0|00"
"0|00|0"
"00|0|0"
*/
func numWays(s string) int {

	mod := int(1e9 + 7)
	cnt := strings.Count(s, "1")

	n := len(s)
	if cnt%3 != 0 {
		return 0
	}
	cnt /= 3
	pre := 0

	cc1 := 0
	cc := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			pre++
		}
		if pre == cnt {
			cc++
		}
		if pre == cnt+1 {
			if cc1 != 0 {
				return cc1 * cc % mod
			}
			pre = 1
			cc1 = cc
			cc = 0
			if pre == cnt {
				cc = 1
			}
		}
	}

	return (n - 1) * (n - 2) / 2 % mod
}

func main() {
	fmt.Println(numWays("00000"))
}
