package main

import (
	"fmt"
	"strings"
)

/*
*
给你一个字符串 word ，返回 word 的所有子字符串中 元音的总数 ，元音是指 'a'、'e'、'i'、'o' 和 'u' 。

子字符串 是字符串中一个连续（非空）的字符序列。

注意：由于对 word 长度的限制比较宽松，答案可能超过有符号 32 位整数的范围。计算时需当心。

输入：word = "aba"
输出：6
解释：
所有子字符串是："a"、"ab"、"aba"、"b"、"ba" 和 "a" 。
- "b" 中有 0 个元音
- "a"、"ab"、"ba" 和 "a" 每个都有 1 个元音
- "aba" 中有 2 个元音
因此，元音总数 = 0 + 1 + 1 + 1 + 1 + 2 = 6 。

思路
1.单独计算每个元音字母可能出现在多少个子串中
2.将所有结果累加
*/
func countVowels(word string) (ans int64) {
	for i, ch := range word {
		if strings.ContainsRune("aeiou", ch) {
			ans += int64(i+1) * int64(len(word)-i)
		}
	}
	return
}

/*
*
dp做法 计算以当前位置结尾，有多少字符，然后再累加

当 i>0 时，以下标 i结尾的所有子字符串中的元音总数根据以下标 i−1 结尾的所有子字符串中的元音总数与  是否为元音计算得到。

著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func countVowels2(word string) (ans int64) {

	n := len(word)
	dp := make([]int, n+1)
	dp[0] = 0
	s := 0

	for i := range word {
		if strings.ContainsRune("aeiou", rune(word[i])) {
			dp[i+1] = dp[i] + i + 1
		} else {
			dp[i+1] = dp[i]
		}
		s += dp[i+1]
	}
	return int64(s)
}
func countVowels3(word string) (ans int64) {

	pre := 0
	s := 0

	for i := range word {
		if strings.ContainsRune("aeiou", rune(word[i])) {
			pre += i + 1
		}
		s += pre
	}
	return int64(s)
}

func main() {
	fmt.Println(countVowels("aba"))
	fmt.Println(countVowels2("aba"))
}
