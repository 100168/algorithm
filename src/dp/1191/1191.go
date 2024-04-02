package main

import "slices"

func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)

	mod := int64(1_000_000_007)
	sum := int64(0)
	for _, v := range arr {
		sum += int64(v)
	}
	dp := make([]int64, 2*n)
	ans := int64(0)
	dp[0] = int64(arr[0])

	for i := 1; i < 2*n; i++ {
		dp[i] = max(dp[i-1]+int64(arr[i%n]), int64(arr[i%n]))
		ans = max(ans, dp[i])
	}
	if k == 1 {
		return int(max(slices.Max(dp[:n]), int64(0)) % mod)
	}
	return int(max(max(int64(k-2)*sum+ans, int64(0)), ans) % mod)

}
