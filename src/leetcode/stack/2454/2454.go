package main

import (
	"fmt"
	"slices"
)

/*
*
给你一个下标从 0 开始的非负整数数组 nums 。对于 nums 中每一个整数，你必须找到对应元素的 第二大 整数。
如果 nums[j] 满足以下条件，那么我们称它为 nums[i] 的 第二大 整数：
j > i
nums[j] > nums[i]
恰好存在 一个 k 满足 i < k < j 且 nums[k] > nums[i] 。
如果不存在 nums[j] ，那么第二大整数为 -1 。

比方说，数组 [1, 2, 4, 3] 中，1 的第二大整数是 4 ，2 的第二大整数是 3 ，3 和 4 的第二大整数是 -1 。
请你返回一个整数数组 answer ，其中 answer[i]是 nums[i] 的第二大整数。
*/
func secondGreaterElement(nums []int) []int {

	n := len(nums)
	ans := make([]int, n)

	for i := range ans {
		ans[i] = -1
	}

	st1 := make([]int, 0)
	st2 := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(st1) > 0 && nums[st1[len(st1)-1]] < nums[i] {
			cur := st1[len(st1)-1]
			st1 = st1[:len(st1)-1]
			ans[cur] = nums[i]

		}
		temp := make([]int, 0)
		for len(st2) > 0 && nums[i] > nums[st2[len(st2)-1]] {
			cur := st2[len(st2)-1]
			temp = append(temp, cur)
			st2 = st2[:len(st2)-1]
		}
		st2 = append(st2, i)
		slices.Reverse(temp)
		st1 = append(st1, temp...)
	}
	return ans
}

func main() {
	fmt.Println(secondGreaterElement([]int{11, 13, 15, 12, 0, 15, 12, 11, 9}))
}
