package main

import "math"

// 也可以用二分
func splitArray(nums []int, k int) int {

	n := len(nums)

	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, k+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i int, j int) int {
		if i < 0 {
			if j == 0 {
				return 0
			}
			return math.MaxInt
		}
		if j == 0 {
			return math.MaxInt
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		cur := math.MaxInt
		sum := 0
		for k := i; k+1 >= j; k-- {
			sum += nums[k]
			cur = min(cur, max(sum, dfs(k-1, j-1)))

		}
		memo[i][j] = cur

		return cur
	}

	return dfs(n-1, k)
}
