package main

func numberOfWays(n int, x int) int {

	var dfs func(int, int) int

	mod := 1_000_000_007

	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dfs = func(i int, rest int) int {
		if rest == 0 {
			return 1
		}
		if rest < 0 || i > n {
			return 0
		}
		if cache[i][rest] != -1 {
			return cache[i][rest]
		}
		cur := dfs(i+1, rest)
		cur += dfs(i+1, rest-power(i, x))
		cache[i][rest] = cur % mod
		return cur % mod

	}
	return dfs(1, n)
}

func power(x, p int) int {
	base := x
	for p > 1 {
		x *= base
		p--
	}
	return x
}

func dp(n int, x int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	mod := 1_000_000_007

	for i := 1; i <= n; i++ {
		for j := n; j >= 0; j-- {
			if j-power(i, x) >= 0 {
				dp[j] += dp[j-power(i, x)]
				dp[j] %= mod
			}
		}
	}
	return dp[n]
}

func main() {
	println(numberOfWays(300, 1))
	println(dp(300, 1))

}
