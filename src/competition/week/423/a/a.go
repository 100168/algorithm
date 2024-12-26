package main

/*
*
给你一个由 n 个整数组成的数组 nums 和一个整数 k，
请你确定是否存在 两个 相邻 且长度为 k 的 严格递增 子数组。
具体来说，需要检查是否存在从下标 a 和 b (a < b) 开始的 两个 子数组，并满足下述全部条件：

这两个子数组 nums[a..a + k - 1] 和 nums[b..b + k - 1] 都是 严格递增 的。
这两个子数组必须是 相邻的，即 b = a + k。
如果可以找到这样的 两个 子数组，请返回 true；否则返回 false。

子数组 是数组中的一个连续 非空 的元素序列。

示例 1：

输入：nums = [2,5,7,8,9,2,3,4,3,1], k = 3

输出：true

解释：

从下标 2 开始的子数组为 [7, 8, 9]，它是严格递增的。
从下标 5 开始的子数组为 [2, 3, 4]，它也是严格递增的。
两个子数组是相邻的，因此结果为 true。
*/
func hasIncreasingSubarrays(nums []int, k int) bool {

	n := len(nums)

	f := make([]int, n)
	f[0] = 1
	for i, v := range nums[1:] {
		if v > nums[i] {
			f[i+1] = f[i] + 1
		} else {
			f[i+1] = 1
		}

		if i+2 >= 2*k {

			if f[i+1] >= k && f[i+1-k] >= k {
				return true
			}
		}
	}
	return false
}
