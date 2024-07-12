package main

import (
	"fmt"
	"slices"
)

/*
*
给你两个下标从 0 开始的数组 nums 和 cost ，分别包含 n 个 正 整数。

你可以执行下面操作 任意 次：

将 nums 中 任意 元素增加或者减小 1 。
对第 i 个元素执行一次操作的开销是 cost[i] 。

请你返回使 nums 中所有元素 相等 的 最少 总开销。

输入：nums = [1,3,5,2], cost = [2,3,1,14]
输出：8
解释：我们可以执行以下操作使所有元素变为 2 ：
- 增加第 0 个元素 1 次，开销为 2 。
- 减小第 1 个元素 1 次，开销为 3 。
- 减小第 2 个元素 3 次，开销为 1 + 1 + 1 = 3 。
总开销为 2 + 3 + 3 = 8 。
这是最小开销。
*/
func minCost(nums []int, cost []int) int64 {

	check := func(m int) int {
		s := 0
		for i, v := range nums {
			diff := v - m
			if diff < 0 {
				diff = -diff
			}
			s += cost[i] * diff
		}
		return s
	}

	l, r := slices.Min(nums), slices.Max(nums)

	for l <= r {
		m := (l + r) / 2
		cur := check(m)
		left := check(m - 1)
		if cur > left {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return int64(min(check(l), check(r)))
}

func main() {
	fmt.Println(minCost([]int{1, 3, 5, 2}, []int{2, 3, 1, 14}))
}
