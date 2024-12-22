package main

/**
给你两个字符串 word1 和 word2 。

如果一个字符串 x 重新排列后，word2 是重排字符串的
前缀
 ，那么我们称字符串 x 是 合法的 。

请你返回 word1 中 合法
子字符串
 的数目。

注意 ，这个问题中的内存限制比其他题目要 小 ，所以你 必须 实现一个线性复杂度的解法。



示例 1：

输入：word1 = "bcca", word2 = "abc"

输出：1

解释：

唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。
*/

func validSubstringCount(word1 string, word2 string) int64 {

	cntC := make([]int, 26)

	for _, v := range word2 {
		cntC[v-'a']++
	}

	n := len(word1)
	l := 0
	ans := 0

	check := func(cnt []int) bool {

		for i, v := range cntC {
			if cnt[i] < v {
				return false
			}
		}
		return true
	}

	cnt := make([]int, 26)

	for i, v := range word1 {
		cnt[v-'a']++
		for check(cnt) {
			cnt[word1[l]-'a']--
			l++
			ans += n - i
		}

	}
	return int64(ans)
}
