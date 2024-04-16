package main

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {

	n := len(nums1)

	nums := [][]int{nums1, nums2}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = []int{1, 1}
	}
	ans := 1
	for i := 1; i < n; i++ {
		for j := 0; j < 2; j++ {
			if nums1[i-1] <= nums[j][i] {
				dp[i][j] = max(dp[i-1][0]+1, dp[i][j])
			}
			if nums2[i-1] <= nums[j][i] {
				dp[i][j] = max(dp[i-1][1]+1, dp[i][j])
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
