package main

import (
	"fmt"
	"math"
)

// n 1
func solution(n int, k int, data []int) int {
	// 初始化 dp 数组
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	// 初始条件
	dp[0][0] = 0
	for j := 1; j <= k; j++ {
		dp[0][j] = data[0] * j
	}

	// 状态转移
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			// 不购买食物
			if j < k {
				dp[i][j] = dp[i-1][j+1]
			}
			// 购买食物
			for m := 1; m <= j; m++ {
				dp[i][j] = min(dp[i][j], dp[i-1][j+1-m]+data[i]*m)
			}
		}
	}

	// 最终结果
	return dp[n-1][1]
}

func main() {
	// Add your test cases here

	fmt.Println(solution(5, 2, []int{1, 2, 3, 3, 2}) == 9)
}
