package main

/*
*
给你一个字符串数组 message 和一个字符串数组 bannedWords。

如果数组中 至少 存在两个单词与 bannedWords 中的任一单词 完全相同，则该数组被视为 垃圾信息。

如果数组 message 是垃圾信息，则返回 true；否则返回 false。

示例 1：

输入： message = ["hello","world","leetcode"], bannedWords = ["world","hello"]

输出： true

解释：

数组 message 中的 "hello" 和 "world" 都出现在数组 bannedWords 中。
*/
func reportSpam(message []string, bannedWords []string) bool {

	banMap := make(map[string]bool)

	for _, v := range bannedWords {
		banMap[v] = true
	}
	cnt := 0
	for _, v := range message {

		if banMap[v] {
			cnt++
		}

	}
	return cnt > 1
}
func main() {

}
