package main

/*
*
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。

一个数组的 能量值 定义为：

如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的
子数组

	的能量值。

请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。

示例 1：

输入：nums = [1,2,3,4,3,2,5], k = 3

输出：[3,4,-1,-1,-1]

解释：

nums 中总共有 5 个长度为 3 的子数组：

[1, 2, 3] 中最大元素为 3 。
[2, 3, 4] 中最大元素为 4 。
[3, 4, 3] 中元素 不是 连续的。
[4, 3, 2] 中元素 不是 上升的。
[3, 2, 5] 中元素 不是 连续的。
*/
func resultsArray(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0)
next:
	for i := k - 1; i < n; i++ {
		for j := i - k + 1; j < i; j++ {
			if nums[j]+1 != nums[j+1] {
				ans = append(ans, -1)
				continue next
			}
		}
		ans = append(ans, nums[i])
	}
	return ans
}