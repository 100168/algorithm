package main

import (
	"fmt"
	"sort"
)

/*
*
给你两个正整数数组 nums 和 target ，两个数组长度相等。

在一次操作中，你可以选择两个 不同 的下标 i 和 j ，其中 0 <= i, j < nums.length ，并且：

令 nums[i] = nums[i] + 2 且
令 nums[j] = nums[j] - 2 。
如果两个数组中每个元素出现的频率相等，我们称两个数组是 相似 的。

请你返回将 nums 变得与 target 相似的最少操作次数。测试数据保证 nums 一定能变得与 target 相似。

示例 1：

输入：nums = [8,12,6], target = [2,14,10]

[6,8,14]
[2,10,14]
输出：2
解释：可以用两步操作将 nums 变得与 target 相似：
- 选择 i = 0 和 j = 2 ，nums = [10,12,4] 。
- 选择 i = 1 和 j = 2 ，nums = [10,14,2] 。
2 次操作是最少需要的操作次数。
*/
func makeSimilar(nums []int, target []int) int64 {

	sort.Ints(nums)
	sort.Ints(target)

	n := len(nums)
	r := n - 1

	add := 0
	sub := 0
	for i := n - 1; i >= 0; i-- {

		if nums[i]%2 == 1 {
			continue
		}
		for r >= 0 && target[r]%2 == 1 {
			r--
		}
		if r < 0 {
			break
		}

		if nums[i] >= target[r] {
			ops := (nums[i] - target[r]) / 2
			add += ops
			sub -= ops
		} else {
			ops := (target[r] - nums[i]) / 2
			sub += ops
			add -= ops
		}
		r--

	}
	r = n - 1
	for i := n - 1; i >= 0; i-- {

		if nums[i]%2 == 0 {
			continue
		}
		for r >= 0 && target[r]%2 == 0 {
			r--
		}
		if r < 0 {
			break
		}

		if nums[i] >= target[r] {
			ops := (nums[i] - target[r]) / 2
			add += ops
			sub -= ops
		} else {
			ops := (target[r] - nums[i]) / 2
			sub += ops
			add -= ops
		}

		r--
	}
	return int64(max(sub, add))

}

// {1, 2, 5}
// {1, 3, 4}
func main() {
	//fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))
}
