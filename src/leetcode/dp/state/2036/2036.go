package main

import "math"

func maximumAlternatingSubarraySum(nums []int) int64 {

	n := len(nums)
	po := make([]int, n)
	ne := make([]int, n)
	//+
	po[0] = nums[0]
	//-
	ne[0] = math.MinInt / 2
	ans := po[0]

	for i := 1; i < n; i++ {
		po[i] = max(ne[i-1], 0) + nums[i]
		//注意-前面必须有+
		ne[i] = po[i-1] - nums[i]
		ans = max(po[i], ne[i], ans)
	}
	return int64(ans)

}
