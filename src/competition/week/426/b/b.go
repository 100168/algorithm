package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个整数数组 nums。该数组包含 n 个元素，其中 恰好 有 n - 2 个元素是 特殊数字 。剩下的 两个 元素中，一个是这些 特殊数字 的 和 ，另一个是 异常值 。

异常值 的定义是：既不是原始特殊数字之一，也不是表示这些数字元素和的数字。

注意，特殊数字、和 以及 异常值 的下标必须 不同 ，但可以共享 相同 的值。

返回 nums 中可能的 最大异常值。

示例 1：

输入： nums = [2,3,5,10]

输出： 10

解释：

特殊数字可以是 2 和 3，因此和为 5，异常值为 10。
*/
func getLargestOutlier(nums []int) int {

	sort.Ints(nums)

	numsMap := make(map[int]int)

	for _, v := range nums {

		numsMap[v]++
	}

	s := 0

	for _, v := range nums {

		s += v
	}

	n := len(nums)

	for i := n - 1; i >= 0; i-- {

		c := s - nums[i]

		if c%2 != 0 {
			continue
		}
		t := c / 2

		if v, ok := numsMap[t]; ok {

			if v > 1 || nums[i] != t {
				return nums[i]
			}

		}

	}

	return 0
}

func main() {
	//fmt.Println(getLargestOutlier([]int{-2, -1, -3, -6, 4}))
	fmt.Println(getLargestOutlier([]int{6, -31, 50, -35, 41, 37, -42, 13}))
}
