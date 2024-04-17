package main

import "slices"

func maxSumAfterOperation(nums []int) int {

	n := len(nums)

	undo := make([]int, n)
	do := make([]int, n)

	//表示不操作
	undo[0] = nums[0]
	//表示操作
	do[0] = nums[0] * nums[0]
	for i := 1; i < n; i++ {
		undo[i] = max(undo[i-1], 0) + nums[i]
		do[i] = max(max(undo[i-1], 0)+(nums[i]*nums[i]), do[i-1]+nums[i])
	}
	return max(slices.Max(undo), slices.Max(do))
}
