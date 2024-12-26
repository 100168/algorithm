package main

/**
给你一个字符串 s 和一个整数 k，在 s 的所有子字符串中，请你统计并返回 至少有一个 字符 至少出现 k 次的子字符串总数。

子字符串 是字符串中的一个连续、 非空 的字符序列。



示例 1：

输入： s = "abacb", k = 2

输出： 4

解释：

符合条件的子字符串如下：

"aba"（字符 'a' 出现 2 次）。
"abac"（字符 'a' 出现 2 次）。
"abacb"（字符 'a' 出现 2 次）。
"bacb"（字符 'b' 出现 2 次）。

*/

func numberOfSubstrings(s string, k int) int {

	ans := 0

	cnt := make([]int, 26)

	l := 0

	check := func() bool {

		for _, v := range cnt {

			if v >= k {
				return true
			}
		}
		return false
	}

	for _, v := range s {

		cnt[v-'a']++

		for check() {
			cnt[s[l]-'a']--
			l++
		}
		ans += l

	}
	return ans
}