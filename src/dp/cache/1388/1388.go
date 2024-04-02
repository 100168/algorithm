package main

import "math"

func maxSizeSlices(slices []int) int {

	n := len(slices)
	m := n / 3
	var dfs func(int, int, int) int

	cache := make([][]int, n)

	for i := range cache {
		cache[i] = make([]int, m+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dfs = func(start, i, rest int) int {
		if i < start {
			if rest == 0 {
				return 0
			}
			return math.MinInt
		}
		if rest == 0 {
			return 0
		}

		if cache[i][rest] != -1 {
			return cache[i][rest]
		}
		//选
		cur := max(slices[i]+dfs(start, i-2, rest-1), dfs(start, i-1, rest))
		cache[i][rest] = cur
		return cur
	}

	ans := dfs(1, n-1, m)
	for i := range cache {
		cache[i] = make([]int, m+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return max(ans, dfs(0, n-2, m))
}

func maxSizeSlices2(slices []int) int {
	n := len(slices)
	return max(dp1(slices, 0, n-2), dp1(slices, 1, n-1))
}

func dp1(nums []int, start int, end int) int {

	n := len(nums)
	m := n / 3
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	dp[start][1] = nums[start]
	dp[start+1][1] = max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-2][j]+nums[i])
		}
	}
	return dp[end][m]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6}
	println(maxSizeSlices(nums))
	println(maxSizeSlices2(nums))
}
