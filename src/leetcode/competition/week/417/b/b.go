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

1.每个元音字母最少出现一次
2.恰好k个非原因字符

思路：
1.
*/
func countOfSubstrings(word string, k int) int {

	cntF := 0
	ans := 0

	n := len(word)
	//有几个连续Z
	pre := make([]int, n)

	cntY := make(map[byte][]int)

	minIndex := -1
	l := -1
	for i, v := range word {
		if strings.ContainsRune("aeiou", v) {
			if l == -1 {
				l = i
			}
			cntY[byte(v)] = append(cntY[byte(v)], i)
			if len(cntY) == 5 && minIndex == -1 {
				minIndex = i
			}
		} else {
			for l <= i {
				cur := cntY[word[l]]
				cur = cur[1:]
				if cur[0] <= minIndex {
					pre[l] = i - minIndex
				}
				cntY[word[l]] = cur
			}
			minIndex = -1
			l = -1
		}
	}

	l = 0

	for i, v := range word {
		if !strings.ContainsRune("aeiou", v) {
			cntF++
		}
		for cntF > k {
			if !strings.ContainsRune("aeiou", rune(word[l])) {
				cntF--
			}
			l++
		}
		if i-l+1-cntF > 0 && cntF == k {
			ans += pre[l]
		}
	}

	return ans
}
func main() {
	fmt.Println(countOfSubstrings("aeioqq", 1))
}
