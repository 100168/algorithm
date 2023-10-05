package main

/*
1525. 字符串的好分割数目
中等
相关标签
相关企业
提示
给你一个字符串 s ，一个分割被称为 「好分割」 当它满足：将 s 分割成 2 个字符串 p 和 q ，它们连接起来等于 s 且 p 和 q 中不同字符的数目相同。

请你返回 s 中好分割的数目。

示例 1：

输入：s = "aacaba"
输出：2
解释：总共有 5 种分割字符串 "aacaba" 的方法，其中 2 种是好分割。
("a", "acaba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
("aa", "caba") 左边字符串和右边字符串分别包含 1 个和 3 个不同的字符。
("aac", "aba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
("aaca", "ba") 左边字符串和右边字符串分别包含 2 个和 2 个不同的字符。这是一个好分割。
("aacab", "a") 左边字符串和右边字符串分别包含 3 个和 1 个不同的字符。
示例 2：

输入：s = "abcd"
输出：1
解释：好分割为将字符串分割成 ("ab", "cd") 。
示例 3：

输入：s = "aaaaa"
输出：4
解释：所有分割都是好分割。
示例 4：

输入：s = "acbadbaada"
输出：2
*/
func numSplits(s string) int {

	cntLeft := make(map[byte]int)
	cntRight := make(map[byte]int)
	n := len(s)
	cntPre := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		cntPre[i] = len(cntRight)
		cntRight[s[i]]++
	}

	ans := 0
	for i, v := range s {
		cntLeft[byte(v)]++
		if len(cntLeft) == cntPre[i] {
			ans++
		}
	}
	return ans

}
