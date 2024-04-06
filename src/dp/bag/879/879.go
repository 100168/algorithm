package main

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {

	mod := 1_000_000_007
	m := len(group)
	sum := 0
	for i := 0; i < m; i++ {
		sum += profit[i]
	}
	cache := make([][][]int, m)
	for i := range cache {
		cache[i] = make([][]int, n+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, sum+1)

			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, p, rest int) int {
		if p < 0 {
			return 0
		}
		if i >= m {
			if rest >= minProfit {
				return 1
			}
			return 0
		}

		if cache[i][p][rest] != -1 {
			return cache[i][p][rest]
		}
		cur := dfs(i+1, p, rest)
		cur += dfs(i+1, p-group[i], rest+profit[i])
		cur %= mod
		cache[i][p][rest] = cur
		return cur
	}
	return dfs(0, n, 0)
}

func dp(n int, minProfit int, group []int, profit []int) int {

	mod := 1_000_000_007
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, minProfit+1)
	}

	dp[0][0] = 1
	for i, v := range group {
		p := profit[i]
		for j := n; j >= v; j-- {
			for k := minProfit; k >= p; k-- {
				dp[j][k] += dp[j-v][max(k-p, 0)]
				dp[j][k] %= mod
			}
		}
	}
	return dp[n-1][minProfit]
}
