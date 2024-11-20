package main

import "slices"

func maximumEnergy(energy []int, k int) int {
	n := len(energy)

	dp := make([]int, n+1)

	dp[0] = 0

	for i := 1; i <= n; i++ {
		dp[i] = energy[i-1]
		if i >= k {
			dp[i] = max(dp[i], dp[i-k]+energy[i-1])
		}
	}

	return slices.Max(dp[n-k+1:])
}
