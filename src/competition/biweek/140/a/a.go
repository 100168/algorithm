package main

import "math"

/*
*

给你一个整数数组 nums 。

请你将 nums 中每一个元素都替换为它的各个数位之 和 。

请你返回替换所有元素以后 nums 中的 最小 元素。

示例 1：

输入：nums = [10,12,13,14]

输出：1

解释：

nums 替换后变为 [1, 3, 4, 5] ，最小元素为 1 。
*/
func minElement(nums []int) int {

	ans := math.MaxInt

	for _, v := range nums {

		cur := 0
		for v > 0 {
			cur += v % 10
			v /= 10
		}
		ans = min(ans, cur)
	}
	return ans

}
