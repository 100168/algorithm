package main

/*
*
一个数组的 最小乘积 定义为这个数组中 最小值 乘以 数组的 和 。

比方说，数组 [3,2,5] （最小值是 2）的最小乘积为 2 * (3+2+5) = 2 * 10 = 20 。
给你一个正整数数组 nums ，请你返回 nums 任意 非空子数组 的最小乘积 的 最大值 。由于答案可能很大，请你返回答案对  109 + 7 取余 的结果。

请注意，最小乘积的最大值考虑的是取余操作 之前 的结果。题目保证最小乘积的最大值在 不取余 的情况下可以用 64 位有符号整数 保存。

子数组 定义为一个数组的 连续 部分。

示例 1：

输入：nums = [1,2,3,2]
输出：14
解释：最小乘积的最大值由子数组 [2,3,2] （最小值是 2）得到。
2 * (2+3+2) = 2 * 7 = 14 。

思路：单调栈+前后缀分解
*/
func maxSumMinProduct(nums []int) int {

	mod := int(1e9 + 7)
	n := len(nums)

	s := make([]int, n+1)
	for i, v := range nums {
		s[i+1] = s[i] + v
	}

	leftMin := make([]int, n)

	for i := range leftMin {
		leftMin[i] = -1
	}
	rightMin := make([]int, n)

	for i := range rightMin {
		rightMin[i] = n
	}

	var st []int
	for i, v := range nums {
		for len(st) > 0 && v < nums[st[len(st)-1]] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			rightMin[last] = i
		}
		st = append(st, i)
	}
	clear(st)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && nums[i] < nums[st[len(st)-1]] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			leftMin[last] = i
		}
		st = append(st, i)

	}

	ans := 0
	for i, v := range nums {
		l := leftMin[i]
		r := rightMin[i]
		ans = max((s[r]-s[l+1])*v, ans)
	}
	return ans % mod

}
