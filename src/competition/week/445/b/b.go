package main

/*
*
给你一个回文字符串 s。

返回 s 的按字典序排列的最小回文排列。

如果一个字符串从前往后和从后往前读都相同，那么这个字符串是一个回文 字符串。

排列是字符串中所有字符的重排。

如果字符串 a 按字典序小于字符串 b，则表示在第一个不同的位置，a 中的字符比 b 中的对应字符在字母表中更靠前。
如果在前 min(a.length, b.length) 个字符中没有区别，则较短的字符串按字典序更小。

示例 1：

输入： s = "z"

输出： "z"

解释：

仅由一个字符组成的字符串已经是按字典序最小的回文。

示例 2：

输入： s = "babab"

输出： "abbba"

解释：

通过重排 "babab" → "abbba"，可以得到按字典序最小的回文。©leetcode
*/
func smallestPalindrome(s string) string {

	cnt := make([]int, 26)

	for _, v := range s {
		cnt[v-'a']++
	}

	ans1 := make([]byte, 0)

	flag := -1

	for i, v := range cnt {

		if v%2 != 0 {
			flag = i
		}

		for j := 0; j < v/2; j++ {
			ans1 = append(ans1, byte(i)+'a')
		}

	}
	c := len(ans1)
	if flag >= 0 {
		ans1 = append(ans1, byte(flag)+'a')
	}

	for i := c - 1; i >= 0; i-- {
		ans1 = append(ans1, byte(i)+'a')
	}

	return string(ans1)

}
