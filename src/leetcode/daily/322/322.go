package main

import "math"

func coinChange(coins []int, amount int) int {

	m := len(coins)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
		dp[i][0] = 0
	}

	var dfs func(int, int) int

	dfs = func(index, rest int) int {
		if rest < 0 {
			return math.MaxInt / 2
		}
		if rest == 0 {
			return 0
		}
		if index == m {
			return math.MaxInt / 2
		}
		if dp[index][rest] != -1 {
			return dp[index][rest]
		}
		cur := math.MaxInt / 2
		for i := 0; i*coins[index] <= rest; i++ {
			cur = min(cur, dfs(index+1, rest-i*coins[index])+i)
		}
		dp[index][rest] = cur
		return cur
	}
	dfs(0, amount)
	if dp[0][amount] == math.MaxInt/2 {
		return -1
	}
	return dp[0][amount]
}
func main() {
	println(coinChange([]int{1, 2, 5}, 11))
}
