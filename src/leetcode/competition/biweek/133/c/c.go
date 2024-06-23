package main

/*
*
给你一个二进制数组 nums 。

你可以对数组执行以下操作 任意 次（也可以 0 次）：

选择数组中 任意 一个下标 i ，并将从下标 i 开始一直到数组末尾 所有 元素 反转 。
反转 一个元素指的是将它的值从 0 变 1 ，或者从 1 变 0 。

请你返回将 nums 中所有元素变为 1 的 最少 操作次数。
*/
func minOperations(nums []int) int {

	n := len(nums)

	ans := 0

	flips := 0

	for i := 0; i < n; i++ {
		if flips%2 == 0 {
			if nums[i] == 0 {
				flips++
				ans++
			}
		} else {
			if nums[i] == 1 {
				flips++
				ans++
			}
		}
	}
	return ans

}
