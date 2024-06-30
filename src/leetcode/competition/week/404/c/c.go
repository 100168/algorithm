package main

import "fmt"

/*
*
给你一个整数数组 nums 和一个 正 整数 k 。
nums 的一个
子序列

	sub 的长度为 x ，如果其满足以下条件，则称其为 有效子序列 ：

(sub[0] + sub[1]) % k == (sub[1] + sub[2]) % k == ... == (sub[x - 2] + sub[x - 1]) % k
返回 nums 的 最长有效子序列 的长度。
*/
func maximumLength(nums []int, k int) int {

	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			cur := (nums[i] + nums[j]) % k
			dp[i][cur] = max(dp[j][cur]+1, dp[i][cur])
			ans = max(ans, dp[i][cur])
		}
	}
	return ans + 1
}

func maximumLength2(nums []int, k int) int {

	n := len(nums)

	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	ans := 0
	for i := 0; i < n; i++ {
		c := nums[i] % k
		for j, v := range dp[c] {
			dp[j][c] = v + 1
			ans = max(ans, dp[j][c])
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 3, 4, 5}, 2))
}
