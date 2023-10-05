package main

func maximumTripletValue2(nums []int) int64 {

	n := len(nums)
	ans := int64(0)
	maxLeft := int64(nums[0])

	dp := make([]int64, n)
	dp[n-1] = int64(nums[n-1])

	for i := n - 2; i >= 0; i-- {
		dp[i] = int64(int(max(int64(nums[i]), dp[i+1])))
	}
	for i := 1; i < n-1; i++ {
		cur := (maxLeft - int64(nums[i])) * dp[i+1]
		ans = max(cur, ans)
		maxLeft = max(maxLeft, int64(nums[i]))
	}
	return ans
}
