package main

/*
*
给你一个字符串 s 。

你需要对 s 执行以下操作 任意 次：

选择一个下标 i ，满足 s[i] 左边和右边都 至少 有一个字符与它相同。
删除 s[i] 左边 离它 最近 且相同的字符。
删除 s[i] 右边 离它 最近 且相同的字符。
请你返回执行完所有操作后， s 的 最短 长度。

示例 1：

输入：s = "abaacbcbb"

输出：5

解释：
我们执行以下操作：

选择下标 2 ，然后删除下标 0 和 3 处的字符，得到 s = "bacbcbb" 。
选择下标 3 ，然后删除下标 0 和 5 处的字符，得到 s = "acbcb" 。
*/
func minimumLength(s string) int {

	indexMap := make(map[rune]int)
	for _, v := range s {
		indexMap[v]++
		if indexMap[v] == 3 {
			indexMap[v] = 1
		}
	}
	ans := 0
	for _, v := range indexMap {
		ans += v
	}
	return ans
}
