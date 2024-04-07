package main

import "strings"

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1
	ans := 0
	for _, v := range strs {
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				zero := strings.Count(v, "0")
				one := len(v) - zero
				if i >= zero && j >= one {
					dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
				}
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
