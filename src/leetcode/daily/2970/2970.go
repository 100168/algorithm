package main

import (
	"fmt"
)

/*
*
给你一个下标从 0 开始的 正 整数数组 nums 。
如果 nums 的一个子数组满足：移除这个子数组后剩余元素 严格递增 ，那么我们称这个子数组为 移除递增 子数组。
比方说，[5, 3, 4, 6, 7] 中的 [3, 4] 是一个移除递增子数组，因为移除该子数组后，[5, 3, 4, 6, 7] 变为 [5, 6, 7] ，是严格递增的。

请你返回 nums 中 移除递增 子数组的总数目。

注意 ，剩余元素为空的数组也视为是递增的。

子数组 指的是一个数组中一段连续的元素序列。

输入：nums = [1,2,3,4]
输出：10
解释：10 个移除递增子数组分别为：[1], [2], [3], [4], [1,2], [2,3], [3,4], [1,2,3], [2,3,4] 和 [1,2,3,4]。移除任意一个子数组后，剩余元素都是递增的。注意，空数组不是移除递增子数组。
示例 2：

输入：nums = [6,5,7,8]
输出：7
解释：7 个移除递增子数组分别为：[5], [6], [5,7], [6,5], [5,7,8], [6,5,7] 和 [6,5,7,8] 。
nums 中只有这 7 个移除递增子数组。
*/
func incremovableSubarrayCount(nums []int) int {

	n := len(nums)
	l, r := 0, n-1

	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			break
		} else {
			l = i
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			break
		} else {
			r = i
		}
	}

	if l == n-1 {
		return n * (n + 1) / 2
	}

	ans := l + 2

	// 6, 5, 7, 8
	//[5], [6], [5,7], [6,5], [5,7,8], [6,5,7] 和 [6,5,7,8] 。
	for i := n - 1; i >= r; i-- {
		for l >= 0 && nums[l] >= nums[i] {
			l--
		}
		ans += l + 2
	}
	return ans
}

func main() {
	fmt.Println(incremovableSubarrayCount([]int{1, 2, 3, 4}))
	fmt.Println(incremovableSubarrayCount([]int{6, 5, 7, 8}))
}
