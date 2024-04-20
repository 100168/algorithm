package main

func wordBreak(s string, wordDict []string) bool {

	n := len(s)

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, v := range wordDict {
			l := len(v)
			if i >= l && dp[i-l] && s[i-l:i] == v {
				dp[i] = true
			}
		}

	}
	return dp[n]
}
