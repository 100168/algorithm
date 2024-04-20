package main

import "math"

func minCut(s string) int {

	n := len(s)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 1; i <= n; i++ {
		dp[i][i] = true
		for j := 1; j < i; j++ {
			if j+1 == i {
				dp[i][j] = s[i-1] == s[j-1]
			} else {
				dp[i][j] = dp[i-1][j+1] && s[i-1] == s[j-1]
			}
		}
	}

	var dfs func(int) int

	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	dfs = func(i int) int {
		if memo[i] != -1 {
			return memo[i]
		}
		if dp[i][1] {
			return 0
		}

		cur := math.MaxInt / 2
		for j := i; j >= 2; j-- {
			if dp[i][j] {
				cur = min(dfs(j-1)+1, cur)
			}
		}
		memo[i] = cur
		return cur
	}

	return dfs(n)

}

func main() {
	println(minCut("abbab"))
}
