package main

func numWays(n int, k int) int {

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	//dp[i][0]表示前i个栅栏颜色不相同的个数 ==>dp[i][0] = dp[i-1][0]*(k-1)+dp[i-1][1]*(k-1)
	//dp[i][1]表示前i个栅栏颜色相同的个数==>dp[i][1] = dp[i-1][0]
	//
	dp[0][0] = k

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]*(k-1) + dp[i-1][1]*(k-1)
		dp[i][1] = dp[i-1][0]
	}
	return dp[n-1][0] + dp[n-1][1]
}
