package main

import (
	"math"
	"slices"
)

func minimumOperations(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, 10)
	}
	for j := 0; j < n; j++ {
		cnt := make([]int, 10)
		for i := 0; i < m; i++ {
			cnt[grid[i][j]]++
		}
		for i, v := range cnt {
			dp[j+1][i] = math.MaxInt
			for k, c := range dp[j] {
				if i == k {
					continue
				}
				dp[j+1][i] = min(dp[j+1][i], c+m-v)
			}
		}

	}
	return slices.Min(dp[n])
}

func main() {
	println(minimumOperations([][]int{{2}}))
}
