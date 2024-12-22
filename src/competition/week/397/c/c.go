package main

import (
	"fmt"
	"math"
)

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt / 2
		}
	}

	// 从右下角开始向上向左遍历
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k < i; k++ {
				dp[i][j] = max(dp[i][j], max(dp[k][j], 0)+grid[i-1][j-1]-grid[k-1][j-1])
			}
			for k := 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], max(dp[i][k], 0)+grid[i-1][j-1]-grid[i-1][k-1])
			}
		}
	}

	// 找到dp中的最大值（即最大得分）
	maxScore := math.MinInt
	for _, row := range dp {
		for _, score := range row {
			if score > maxScore {
				maxScore = score
			}
		}
	}

	// 题目要求至少移动一次，因此最大得分可能不包括起点的值
	// 但由于我们已经从右下角开始计算，所以dp中的最大值已经是考虑过的最大得分
	return maxScore
}

func main() {
	grid := [][]int{
		{9, 5, 7, 3},
		{8, 9, 6, 1},
		{6, 7, 14, 3},
		{2, 5, 3, 1}}

	fmt.Println(maxScore(grid)) // 输出应为 9
}
