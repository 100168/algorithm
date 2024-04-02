package main

func rob(nums []int) int {

	n := len(nums)
	cache := make([]int, n)

	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 0
		}

		if cache[i] != -1 {
			return cache[i]
		}

		cur := max(nums[i]+dfs(i-2), dfs(i-1))

		cache[i] = cur
		return cur

	}
	return dfs(n - 1)
}

func rob2(nums []int) int {

	n := len(nums)
	dp := make([]int, n+1)

	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(nums[i-1]+dp[i-2], dp[i-1])
	}
	return dp[n]
}
