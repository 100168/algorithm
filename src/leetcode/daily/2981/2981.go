package main

import (
	"fmt"
)

/*
给你一个仅由小写英文字母组成的字符串 s 。

如果一个字符串仅由单一字符组成，那么它被称为 特殊 字符串。例如，字符串 "abc" 不是特殊字符串，而字符串 "ddd"、"zz" 和 "f" 是特殊字符串。

返回在 s 中出现 至少三次 的 最长特殊子字符串 的长度，如果不存在出现至少三次的特殊子字符串，则返回 -1 。

子字符串 是字符串中的一个连续 非空 字符序列。
*/
func maximumLength(s string) int {

	pre := 0
	cnt := make(map[string]int)
	ans := -1
	for i := 0; i < len(s); i++ {
		if s[i] != s[pre] {
			pre = i
		}
		for j := pre + 1; j <= i+1; j++ {
			cnt[s[pre:j]]++
			if cnt[s[pre:j]] >= 3 {
				ans = max(ans, j-pre)
			}
		}

	}

	return ans
}

func maximumLength2(s string) int {
	ans := -1
	length := len(s)

	chs := make([][]int, 26)
	cnt := 0
	for i := 0; i < length; i++ {
		cnt++
		if i+1 == length || s[i] != s[i+1] {
			ch := int(s[i] - 'a')
			chs[ch] = append(chs[ch], cnt)
			cnt = 0

			for j := len(chs[ch]) - 1; j > 0; j-- {
				if chs[ch][j] > chs[ch][j-1] {
					tmp := chs[ch][j-1]
					chs[ch][j-1] = chs[ch][j]
					chs[ch][j] = tmp
				} else {
					break
				}
			}

			if len(chs[ch]) > 3 {
				chs[ch] = chs[ch][:len(chs[ch])-1]
			}
		}
	}

	for i := 0; i < 26; i++ {
		if len(chs[i]) > 0 && chs[i][0] > 2 {
			ans = max(ans, chs[i][0]-2)
		}
		if len(chs[i]) > 1 && chs[i][0] > 1 {
			ans = max(ans, min(chs[i][0]-1, chs[i][1]))
		}
		if len(chs[i]) > 2 {
			ans = max(ans, chs[i][2])
		}
	}

	return ans
}

func main() {
	fmt.Println(maximumLength("cccerrrecdcdccedecdcccddeeeddcdcddedccdceeedccecde"))
}
