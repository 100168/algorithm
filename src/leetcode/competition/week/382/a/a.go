package main

/*
*
给你一个下标从 0 开始的字符串 s ，该字符串由用户输入。按键变更的定义是：使用与上次使用的按键不同的键。
例如 s = "ab" 表示按键变更一次，而 s = "bBBb" 不存在按键变更。

返回用户输入过程中按键变更的次数。

注意：shift 或 caps lock 等修饰键不计入按键变更，也就是说，如果用户先输入字母 'a' 然后输入字母 'A' ，不算作按键变更。

示例 1：

输入：s = "aAbBcC"
输出：2
解释：
从 s[0] = 'a' 到 s[1] = 'A'，不存在按键变更，因为不计入 caps lock 或 shift 。
从 s[1] = 'A' 到 s[2] = 'b'，按键变更。
从 s[2] = 'b' 到 s[3] = 'B'，不存在按键变更，因为不计入 caps lock 或 shift 。
从 s[3] = 'B' 到 s[4] = 'c'，按键变更。
从 s[4] = 'c' 到 s[5] = 'C'，不存在按键变更，因为不计入 caps lock 或 shift 。
*/
func countKeyChanges(s string) int {

	ans := 0

	for i := range s[1:] {
		if s[i]|32 != s[i+1]|32 {
			ans++
		}
	}
	return ans
}
