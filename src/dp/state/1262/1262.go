package main

import (
	"math"
)

func maxSumDivThree(nums []int) int {

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	dp[0][1] = math.MinInt
	dp[0][2] = math.MinInt
	for i, x := range nums {
		for j := 0; j < 3; j++ {
			dp[(i+1)%2][j] = max(dp[i%2][j], dp[i%2][((j-x)%3+3)%3]+x)
		}
	}

	m := len(nums)
	return dp[m%2][0]

}
