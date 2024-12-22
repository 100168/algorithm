package main

import (
	"strings"
)

/*
给你一个字符串 s 和一个字符 c 。返回在字符串 s 中并且以 c 字符开头和结尾的非空子字符串的总数。

示例 1：

输入：s = "abada", c = "a"

输出：6

解释：以 "a" 开头和结尾的子字符串有： "abada"、"abada"、"abada"、"abada"、"abada"、"abada"。

示例 2：

输入：s = "zzz", c = "z"

输出：6

解释：字符串 s 中总共有 6 个子字符串，并且它们都以 "z" 开头和结尾。
*/
func countSubstrings(s string, c byte) int64 {

	count := strings.Count(s, string(c))
	if count <= 1 {
		return int64(count)
	}

	ans := int64(1)
	ans = int64(count) * int64(count-1)
	return ans/int64(2) + int64(count)
}
