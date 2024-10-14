package main

import (
	"fmt"
	"strings"
)

/*
*
给你一个字符串 word 和一个 非负 整数 k。

返回 word 的
子字符串

	中，每个元音字母（'a'、'e'、'i'、'o'、'u'）至少 出现一次，并且 恰好 包含 k 个辅音字母的子字符串的总数。

示例 1：

输入：word = "aeioqq", k = 1

输出：0

解释：

不存在包含所有元音字母的子字符串。

示例 2：

输入：word = "aeiou", k = 0

输出：1

解释：

唯一一个包含所有元音字母且不含辅音字母的子字符串是 word[0..4]，即 "aeiou"。

示例 3：

输入：word = "ieaouqqieaouqq", k = 1

输出：3

解释：

包含所有元音字母并且恰好含有一个辅音字母的子字符串有：

word[0..5]，即 "ieaouq"。
word[6..11]，即 "qieaou"。
word[7..12]，即 "ieaouq"。

提示：

5 <= word.length <= 250
word 仅由小写英文字母组成。
0 <= k <= word.length - 5
*/
func countOfSubstrings(word string, k int) int {

	cnt := 0

	ans := 0

	n := len(word)

	cntY := make(map[byte]int)
	for i, v := range word {
		if !strings.ContainsRune("aeiou", v) {
			cnt++
		} else {
			cntY[byte(v)]++
		}
		for cnt > k {
			if !strings.ContainsRune("aeiou", rune(word[l])) {
				cnt--
			} else {
				cntY[word[l]]--
				if cntY[word[l]] == 0 {
					delete(cntY, word[l])
				}
			}
			l++
		}
		if cnt == k && len(cntY) > 0 {
			ans++
		}
	}

	return ans
}
func main() {
	fmt.Println(countOfSubstrings("aeioqoq", 1))
}
