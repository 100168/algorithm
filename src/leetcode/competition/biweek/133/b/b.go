package main

/*
*
给你一个二进制数组 nums 。

你可以对数组执行以下操作 任意 次（也可以 0 次）：

选择数组中 任意连续 3 个元素，并将它们 全部反转 。
反转 一个元素指的是将它的值从 0 变 1 ，或者从 1 变 0 。

请你返回将 nums 中所有元素变为 1 的 最少 操作次数。如果无法全部变成 1 ，返回 -1 。
*/
func minOperations(nums []int) int {

	n := len(nums)

	ans := 0

	for i := 0; i < n-3; i++ {
		if nums[i] == 0 {
			ans++
			nums[i+1] ^= 1
			nums[i+2] ^= 1

		}
	}

	if nums[n-1] == 0 || nums[n-2] == 0 {
		return -1
	}
	return ans

}
