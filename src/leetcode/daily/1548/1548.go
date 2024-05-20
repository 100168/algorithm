package main

import "fmt"

/*
*
给你一个字符串 s 。请返回 s 中最长的 超赞子字符串 的长度。

「超赞子字符串」需满足满足下述两个条件：

该字符串是 s 的一个非空子字符串
进行任意次数的字符交换后，该字符串可以变成一个回文字符串

示例 1：

输入：s = "3242415"
输出：5
解释："24241" 是最长地超赞子字符串，交换其中的字符后，可以得到回文 "24142"
示例 2：

输入：s = "12345678"
输出：1
示例 3：

输入：s = "213123"
输出：6
解释："213123" 是最长地超赞子字符串，交换其中的字符后，可以得到回文 "231132"
示例 4：

输入：s = "00"
输出：2

提示：

1 <= s.length <= 10^5
s 仅由数字组成
*/
func longestAwesome(s string) int {

	n := len(s)
	ans := 0
	xorMap := make(map[int]int)
	xorMap[0] = -1
	pre := 0
	for i := 0; i < n; i++ {
		pre ^= 1 << int(s[i]-'0')
		if _, ok := xorMap[pre]; !ok {
			xorMap[pre] = i
		}
		if _, ok := xorMap[pre]; ok {
			ans = max(ans, i-xorMap[pre])
		}
		for j := 0; j <= 10; j++ {
			cur := (1 << j) ^ pre
			if _, ok := xorMap[cur]; ok {
				ans = max(ans, i-xorMap[cur])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestAwesome("3242415"))
}
