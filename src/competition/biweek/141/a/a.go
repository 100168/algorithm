package main

import (
	"math"
)

/*
*
给你一个长度为 n 的
质数
数组 nums 。你的任务是返回一个长度为 n 的数组 ans ，对于每个下标 i ，以下 条件 均成立：

ans[i] OR (ans[i] + 1) == nums[i]
除此以外，你需要 最小化 结果数组里每一个 ans[i] 。

如果没法找到符合 条件 的 ans[i] ，那么 ans[i] = -1 。

质数 指的是一个大于 1 的自然数，且它只有 1 和自己两个因数。

示例 1：

输入：nums = [2,3,5,7]

输出：[-1,1,4,3]
*/
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
