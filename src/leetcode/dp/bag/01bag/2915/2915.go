package main

import (
	"math"
)

func lengthOfLongestSubsequence(nums []int, target int) int {

	m := len(nums)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, target+1)
	}
	var dfs func(int, int) int
	dfs = func(i, rest int) int {
		if rest == 0 {
			return 0
		}
		if i < 0 || rest < 0 {
			return math.MinInt / 2
		}
		if cache[i][rest] != 0 {
			return cache[i][rest]
		}
		cache[i][rest] = max(dfs(i-1, rest), dfs(i-1, rest-nums[i])+1)
		return cache[i][rest]
	}
	ans := dfs(m-1, target)
	if ans < 0 {
		return -1
	}
	return ans
}

func lengthOfLongestSubsequence2(nums []int, target int) int {

	m := len(nums)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := 0; j <= target; j++ {
			dp[i][j] = math.MinInt / 2
		}
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-nums[i-1]]+1)
			}
		}
	}

	if dp[m][target] <= 0 {
		return -1
	}

	return dp[m][target]
}

func lengthOfLongestSubsequence3(nums []int, target int) int {
	m := len(nums)
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = math.MinInt / 2
	}
	dp[0] = 0
	for i := 0; i < m; i++ {
		for j := target; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j], dp[j-nums[i]]+1)
			}
		}
	}

	if dp[target] <= 0 {
		return -1
	}

	return dp[target]
}
func main() {
	println(lengthOfLongestSubsequence2([]int{14, 6, 4, 15, 5, 3, 6, 7, 8}, 36))
	println(lengthOfLongestSubsequence([]int{14, 6, 4, 15, 5, 3, 6, 7, 8}, 36))
}
