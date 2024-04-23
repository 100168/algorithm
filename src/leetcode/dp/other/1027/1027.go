package main

import (
	"fmt"
	"slices"
)

/*
给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。

回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。
并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。
*/
func longestArithSeqLength(nums []int) int {

	n := len(nums)

	numsMap := make(map[int]int)

	maxVal := slices.Max(nums)
	dp := make([][][]int, n)

	for i := range dp {
		dp[i] = make([][]int, maxVal+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = 1
			}
		}
	}

	ans := 1
	for i := 0; i < n; i++ {
		cur := nums[i]
		for j := 0; j <= maxVal; j++ {
			if _, ok := numsMap[cur-j]; ok {
				dp[i][j][0] = dp[numsMap[cur-j]][j][0] + 1
			}
			if _, ok := numsMap[cur+j]; ok {
				dp[i][j][1] = dp[numsMap[cur+j]][j][1] + 1
			}
			ans = max(ans, dp[i][j][0], dp[i][j][1])
		}
		numsMap[cur] = i
	}
	return ans
}

func main() {
	fmt.Println(longestArithSeqLength([]int{22, 15, 8, 1}))

}
