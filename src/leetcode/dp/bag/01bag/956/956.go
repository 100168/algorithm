package main

import "math"

func tallestBillboard(rods []int) int {

	n := len(rods)

	s := 0
	for _, v := range rods {
		s += v
	}

	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, s+1)
		for j := 0; j <= s; j++ {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, diff int) int {
		if i < 0 {
			if diff == 0 {
				return 0
			}
			return math.MinInt
		}
		if cache[i][diff] != -1 {
			return cache[i][diff]
		}
		cur := dfs(i-1, diff)
		cur = max(dfs(i-1, abs(diff-rods[i]))+max(0, rods[i]-diff), cur)
		cur = max(dfs(i-1, diff+rods[i])+rods[i], cur)
		cache[i][diff] = cur
		return cur
	}
	return dfs(n-1, 0)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dp(rods []int) int {

	n := len(rods)

	s := 0
	for _, v := range rods {
		s += v
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, s+1)

		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}

	for i, v := range rods {
		for j := 0; j <= s; j-- {
			dp[i][j] = dp[i-1][j]
			dp[i][j] = max(dp[i-1][abs(j-v)]+max(0, j-v), dp[i][j])
			if j+v <= s {
				dp[i][j] = max(dp[i][j], dp[i-1][j+v]+v)
			}
		}
	}
	return dp[n-1][0]
}
