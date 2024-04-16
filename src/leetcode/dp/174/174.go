package main

import "math"

func calculateMinimumHP(dungeon [][]int) int {

	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = min(dungeon[m-1][n-1], 0)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			dp[i][j] = math.MinInt
			if i < m-1 {
				dp[i][j] = max(dp[i][j], dp[i+1][j])
			}
			if j < n-1 {
				dp[i][j] = max(dp[i][j], dp[i][j+1])
			}

			dp[i][j] += dungeon[i][j]

			if i == 0 && j == 0 {
				break
			}
			dp[i][j] = min(dp[i][j], 0)

		}
	}

	if dp[0][0] <= 0 {
		return -dp[0][0] + 1
	}
	return 1
}

func main() {
	println(calculateMinimumHP([][]int{{-3, 5}}))
}
