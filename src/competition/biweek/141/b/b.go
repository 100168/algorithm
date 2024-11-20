package main

import "math"

func minBitwiseArray(nums []int) []int {

	// nums[i]   nums[i]+1 = nums[i]

	// 101
	// x+1 = nums[i]

	for i, v := range nums {
		cur := math.MaxInt
		for sub := v; sub > 0; sub = (sub - 1) & v {
			if sub|(sub+1) == v {
				cur = min(cur, sub)
			}
		}
		nums[i] = cur
		if nums[i] == math.MaxInt {
			nums[i] = -1
		}

	}
	return nums
}
