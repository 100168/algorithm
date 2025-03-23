package main

import "fmt"

/*
*
给你一个长度为 n 的二进制字符串 s 和一个整数 numOps。

你可以对 s 执行以下操作，最多 numOps 次：

选择任意下标 i（其中 0 <= i < n），并 翻转 s[i]，即如果 s[i] == '1'，则将 s[i] 改为 '0'，反之亦然。
你需要 最小化 s 的最长 相同
子字符串的长度，相同子字符串 是指子字符串中的所有字符都 相同。

返回执行所有操作后可获得的 最小 长度。

示例 1：

输入: s = "000001", numOps = 1

输出: 2

解释:

将 s[2] 改为 '1'，s 变为 "001001"。最长的所有字符相同的子串为 s[0..1] 和 s[3..4]。
*/
func minLength(s string, numOps int) int {

	n := len(s)

	l, r := 1, n

	check := func(t int) bool {

		if t == 1 {

			cnt := 0

			//开始是0
			for i := 0; i < n; i++ {

				cnt += int(s[i]-'0') ^ (i & 1)

			}
			return min(cnt, n-cnt) <= numOps

		}

		cnt := 0
		k := 0
		for i := range s {
			k++

			if i == n-1 || s[i] != s[i+1] {
				cnt += k / (t + 1)
				k = 0
			}

		}
		return cnt <= numOps

	}

	for l <= r {
		m := (l + r) / 2

		if check(m) {

			r = m - 1
		} else {
			l = m + 1
		}

	}
	return l
}
func main() {
	fmt.Println(minLength("0101", 0))
}
