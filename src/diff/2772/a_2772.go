package main

/*
给你一个下标从 0 开始的整数数组 nums 和一个正整数 k 。

你可以对数组执行下述操作 任意次 ：

从数组中选出长度为 k 的 任一 子数组，并将子数组中每个元素都 减去 1 。
如果你可以使数组中的所有元素都等于 0 ，返回  true ；否则，返回 false 。

子数组 是数组中的一个非空连续元素序列。

示例 1：

输入：nums = [2,2,3,1,1,0], k = 3
输出：true
解释：可以执行下述操作：
- 选出子数组 [2,2,3] ，执行操作后，数组变为 nums = [1,1,2,1,1,0] 。
- 选出子数组 [2,1,1] ，执行操作后，数组变为 nums = [1,1,1,0,0,0] 。
- 选出子数组 [1,1,1] ，执行操作后，数组变为 nums = [0,0,0,0,0,0] 。
示例 2：

输入：nums = [1,3,1,1], k = 2
输出：false
解释：无法使数组中的所有元素等于 0 。
*/
func checkArray(nums []int, k int) bool {

	n := len(nums)
	diff := make([]int, n+1)
	diff[0] = nums[0]
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	for i := 0; i+k < n; i++ {
		curSub := diff[i]
		diff[i] = 0
		diff[i+k] += curSub
	}

	for i := 0; i < n; i++ {
		if diff[i] != 0 {
			return false
		}
	}
	return true
}
