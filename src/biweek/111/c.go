package main

func minimumOperations(nums []int) int {

	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 4)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j < 4; j++ {

			dp[i][j] = n
			for k := 1; k <= j; k++ {
				dp[i][j] = min(dp[i-1][k], dp[i][j])
			}
			if nums[i-1] != j {
				dp[i][j]++
			}
		}
	}
	ans := n

	for i := 1; i < 4; i++ {
		if dp[n][i] < ans {
			ans = dp[n][i]
		}
	}

	return ans

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nums := []int{2, 2, 2, 2, 3, 3}
	println(minimumOperations(nums))
}
