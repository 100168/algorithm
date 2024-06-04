package main

import "fmt"

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
*/
func countVowels(word string) int64 {

	ans := 0
	oMap := map[uint8]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	cnt := 0
	for i := 0; i < len(word); i++ {
		if oMap[word[i]] {
			cnt++
		}
	}
	return int64(ans)
}

func main() {
	fmt.Println(countVowels("aba"))
}
