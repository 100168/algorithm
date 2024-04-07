package main

func change(amount int, coins []int) int {

	n := len(coins)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(index, rest int) int {

		if rest < 0 {
			return 0
		}

		if index < 0 {
			if rest == 0 {
				return 1
			}
			return 0
		}

		if dp[index][rest] != -1 {
			return dp[index][rest]
		}

		cur := 0
		cur = dfs(index, rest-coins[index]) + dfs(index-1, rest)
		dp[index][rest] = cur
		return cur
	}
	return dfs(n-1, amount)
}

func main() {
	println(change(5, []int{1, 2, 5}))
}
