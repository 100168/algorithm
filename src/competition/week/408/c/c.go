package main

import (
	"fmt"
)

/*

给你一个二进制字符串 s。

请你统计并返回其中 1 显著 的子字符串的数量。

如果字符串中 1 的数量 大于或等于 0 的数量的 平方，则认为该字符串是一个 1 显著 的字符串 。

示例 1：

输入：s = "00011"

输出：5

解释：

1 显著的子字符串如下表所示。

输入：s = "101101"

输出：16

解释：

1 不显著的子字符串如下表所示。

总共有 21 个子字符串，其中 5 个是 1 不显著字符串，因此有 16 个 1 显著子字符串。

*/

/*
1.枚举左端点
*/
func numberOfSubstrings(s string) (ans int) {
	var a []int
	for i, b := range s {
		if b == '0' {
			a = append(a, i)
		}
	}
	n := len(s)
	tot1 := n - len(a)
	a = append(a, n) // 哨兵

	for left, b := range s {
		if b == '1' {
			ans += a[0] - left // 不含 0 的子串个数
		}
		for k, j := range a[:len(a)-1] {
			cnt0 := k + 1
			if cnt0*cnt0 > tot1 {
				break
			}
			cnt1 := j - left - k
			ans += max(a[k+1]-j-max(cnt0*cnt0-cnt1, 0), 0)
		}
		if b == '0' {
			a = a[1:] // 这个 0 后面不会再枚举到了
		}
	}
	return
}

// / 111111
func main() {
	//fmt.Println(numberOfSubstrings("00011"))
	fmt.Println(numberOfSubstrings("101101"))
}
