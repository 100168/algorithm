package main

import (
	"fmt"
	"slices"
)

/*
*
给你两个字符串 s 和 pattern 。

如果一个字符串 x 修改 至多 一个字符会变成 y ，那么我们称它与 y 几乎相等 。

请你返回 s 中下标 最小 的
子字符串，它与 pattern 几乎相等 。如果不存在，返回 -1 。

子字符串 是字符串中的一个 非空、连续的字符序列。

示例 1：

输入：s = "abcdefg", pattern = "bcdffg"

输出：1

解释：

将子字符串 s[1..6] == "bcdefg" 中 s[4] 变为 "f" ，得到 "bcdffg" 。

示例 2：

输入：s = "ababbababa", pattern = "bacaba"

输出：4

解释：

将子字符串 s[4..9] == "bababa" 中 s[6] 变为 "c" ，得到 "bacaba" 。

示例 3：

输入：s = "abcd", pattern = "dba"

输出：-1

示例 4：

输入：s = "dde", pattern = "d"

输出：0

提示：

1 <= pattern.length < s.length <= 3 * 105
s 和 pattern 都只包含小写英文字母。

进阶：如果题目变为 至多 k 个 连续 字符可以被修改，你可以想出解法吗？

z func
*/
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)

	for i, c, r := 1, 1, 1; i < n; i++ {

		l := min(z[i-c], max(r-i, 0))

		for i+l < n && s[i+l] == s[l] {
			l++
		}
		z[i] = l

		if i+l > r {
			c = i
			r = i + l
		}
	}
	return z
}

func rev(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}

func minStartingIndex(s, pattern string) int {
	preZ := calcZ(pattern + s)
	sufZ := calcZ(rev(pattern) + rev(s))
	slices.Reverse(sufZ) // 也可以不反转，下面写 sufZ[len(sufZ)-i]
	m := len(pattern)
	for i := m; i <= len(s); i++ {
		if preZ[i]+sufZ[i-1] >= m-1 {
			return i - m
		}
	}
	return -1
}

func main() {
	fmt.Println(minStartingIndex("abcdefg", "bcdffg"))
	//fmt.Println(minStartingIndex("ede", "d"))
	//fmt.Println(minStartingIndex("efeff", "fe"))
	//fmt.Println(minStartingIndex("dcdccdddcdcdddccddccccddcd", "ddc"))
	//fmt.Println(minStartingIndex("mnmm", "mm"))
	//fmt.Println(minStartingIndex("cbccbcccb", "bcb"))
	//fmt.Println(minStartingIndex("deeeeddde", "eddd"))
	//fmt.Println(minStartingIndex("ccccbcbbbbbbb", "ccbcb"))
	//fmt.Println(minStartingIndex("eddeeded", "ddede"))

}
