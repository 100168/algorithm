package main

import (
	"fmt"
	"strings"
)

/**
给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。


输入：s = "eleetminicoworoep"
输出：13
解释：最长子字符串是 "leetminicowor" ，它包含 e，i，o 各 2 个，以及 0 个 a，u 。

*/

func findTheLongestSubstring(s string) int {

	bts := 0

	ans := 0
	sum := make(map[int]int)
	sum[0] = -1
	for i, v := range s {
		if strings.ContainsRune("aeiou", v) {
			bts ^= 1 << int(v-'a')
		}
		if _, ok := sum[bts]; ok {
			pre := sum[bts]
			ans = max(ans, i-pre)
		} else {
			sum[bts] = i
		}
	}
	return ans
}

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}
