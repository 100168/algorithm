package main

import "strings"

/*
1653. 使字符串平衡的最少删除次数
中等
相关标签
相关企业
提示
给你一个字符串 s ，它仅包含字符 'a' 和 'b'。

你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a' ，此时认为 s 是 平衡 的。

请你返回使 s 平衡 的 最少 删除次数。

示例 1：

输入：s = "aababbab"
输出：2
解释：你可以选择以下任意一种方案：
下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
示例 2：

输入：s = "bbaaaaabb"
输出：2
解释：唯一的最优解是删除最前面两个字符。

1.dp
2. 贪心+二分
3.字符串前缀
*/
func minimumDeletions(s string) int {
	maxL := 0
	n := len(s)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if s[i] >= s[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxL = max(dp[i], maxL)
	}
	return n - maxL
}

func minimumDeletions2(s string) int {
	n := len(s)
	dp := make([]uint8, n+1)
	l := 1
	dp[l] = s[0]
	for i := 1; i < n; i++ {
		if dp[l] <= s[i] {
			dp[l+1] = s[i]
			l++
		} else {
			left, right, pos := 1, l, 0
			for left <= right {
				mid := (left + right) / 2
				if dp[mid] <= s[i] {
					pos = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			dp[pos+1] = s[i]
		}
	}
	return n - l
}
func minimumDeletions3(s string) int {
	del := strings.Count(s, "a")

	ans := del
	for _, v := range s {
		del += int(v-'a')*2 - 1
		ans = min(ans, del)
	}
	return ans
}
