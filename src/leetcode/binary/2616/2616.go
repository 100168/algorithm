package main

import (
	"slices"
	"sort"
)

func minimizeMax(nums []int, p int) int {

	sort.Ints(nums)

	l, r := 0, slices.Max(nums)

	for l <= r {
		m := (r + l) / 2
		if check(nums, m, p) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(nums []int, t, p int) bool {

	n := len(nums)

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] <= t {
			if i >= 2 {
				dp[i] = dp[i-2] + 1
			} else {
				dp[i] = 1
			}
		}
		dp[i] = max(dp[i-1], dp[i])

	}
	return dp[n-1] >= p

}
