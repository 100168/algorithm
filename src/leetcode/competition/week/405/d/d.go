package main

/**
给你一个字符串 target、一个字符串数组 words 以及一个整数数组 costs，这两个数组长度相同。

设想一个空字符串 s。

你可以执行以下操作任意次数（包括零次）：

选择一个在范围  [0, words.length - 1] 的索引 i。
将 words[i] 追加到 s。
该操作的成本是 costs[i]。
返回使 s 等于 target 的 最小 成本。如果不可能，返回 -1。
*/

func minimumCost(target string, words []string, costs []int) int {

	n := len(target)

	wordMap := make(map[string]int)

	maxL := 0
	for i, w := range words {
		if _, ok := wordMap[w]; !ok {
			wordMap[w] = costs[i]
		} else {
			wordMap[w] = min(costs[i], wordMap[w])
		}
		maxL = max(maxL, len(w))
	}

	dp := make([]int, n)

	for i := 1; i <= n; i++ {

		for j := i - 1; j >= 0 && i-j <= maxL; j-- {
			cur := target[j:i]
			if _, ok := wordMap[cur]; ok {
				dp[i] = min(dp[i], dp[i-j]+wordMap[cur])
			}
		}

	}

	return dp[n]
}
