package main

func checkPartitioning(s string) bool {

	n := len(s)

	memo := make([]map[int]bool, n+1)

	for i := range memo {
		memo[i] = make(map[int]bool)
	}
	dp := make([][]bool, n+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	for i := 1; i <= n; i++ {

		dp[i][i] = true

		for j := 1; j < i; j++ {

			if j+1 == i {
				if s[j-1] == s[i-1] {
					dp[i][j] = true
				}
			} else {
				dp[i][j] = dp[i-1][j+1] && s[i-1] == s[j-1]
			}
		}
	}

	var dfs func(int, int) bool

	dfs = func(i int, rest int) bool {
		if i == 0 {
			return rest == 0
		}

		if rest <= 0 {
			return false
		}
		if _, ok := memo[i][rest]; ok {
			return memo[i][rest]
		}

		cur := false

		for j := 1; j <= i; j++ {

			if dp[i][j] {
				cur = cur || dfs(j-1, rest-1)
			}
		}

		memo[i][rest] = cur

		return cur

	}
	return dfs(n, 3)
}
