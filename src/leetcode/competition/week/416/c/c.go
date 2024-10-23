package main

import "fmt"

/*
*
给你两个字符串 word1 和 word2 。

如果一个字符串 x 重新排列后，word2 是重排字符串的
前缀
 ，那么我们称字符串 x 是 合法的 。

请你返回 word1 中 合法
子字符串
 的数目。



示例 1：

输入：word1 = "bcca", word2 = "abc"

输出：1

解释：

唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。


*/

func validSubstringCount(word1 string, word2 string) int64 {

	cnt := make([]int, 26)
	for _, v := range word2 {
		cnt[v-'a']++
	}

	cntC := make([]int, 26)

	check := func() bool {
		for i, v := range cnt {

			if v > cntC[i] {
				return false
			}
		}
		return true
	}

	//l左边都是满足条件
	l := 0
	ans := 0
	for _, v := range word1 {
		cntC[v-'a']++
		for check() {
			cntC[word1[l]-'a']--
			l++
		}
		ans += l
	}
	return int64(ans)
}
func main() {

	fmt.Println(validSubstringCount("bcca", "abc"))

}
