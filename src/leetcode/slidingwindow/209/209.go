package main

import "math"

/*
*给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 连续
子数组

	[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	l := 0

	s := 0

	ans := math.MaxInt
	for r := 0; r < n; r++ {
		s += nums[r]
		for s >= target {
			ans = min(ans, r-l+1)
			s -= nums[l]
			l++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}
