package main

import (
	"fmt"
	"math"
)

/*
*
给你一个长度为 n、下标从 0 开始的整数数组 nums，nums[i] 表示收集位于下标 i 处的巧克力成本。
每个巧克力都对应一个不同的类型，最初，位于下标 i 的巧克力就对应第 i 个类型。

在一步操作中，你可以用成本 x 执行下述行为：

同时修改所有巧克力的类型，将巧克力的类型 ith 修改为类型 ((i + 1) mod n)th。
假设你可以执行任意次操作，请返回收集所有类型巧克力所需的最小成本。

输入：nums = [20,1,15], x = 5
输出：13
解释：最开始，巧克力的类型分别是 [0,1,2] 。我们可以用成本 1 购买第 1 个类型的巧克力。
接着，我们用成本 5 执行一次操作，巧克力的类型变更为 [1,2,0] 。我们可以用成本 1 购买第 2 个类型的巧克力。
然后，我们用成本 5 执行一次操作，巧克力的类型变更为 [2,0,1] 。我们可以用成本 1 购买第 0 个类型的巧克力。
因此，收集所有类型的巧克力需要的总成本是 (1 + 5 + 1 + 5 + 1) = 13 。可以证明这是一种最优方案。

输入：nums = [1,2,3], x = 4
输出：6
解释：我们将会按最初的成本收集全部三个类型的巧克力，而不需执行任何操作。因此，收集所有类型的巧克力需要的总成本是 1 + 2 + 3 = 6 。

思路：枚举
1.枚举n-1次
2.使用一个数组，记录当前枚举范围内的最小值
*/
func minCost(nums []int, x int) int64 {

	n := len(nums)
	minV := make([]int, n)
	for i := range nums {
		minV[i] = nums[i]
	}
	ans := math.MaxInt
	c := 0
	for ops := 0; ops < n; ops++ {
		cur := c
		for i := range nums {
			minV[i] = min(minV[i], nums[(ops+i)%n])
			cur += minV[i]
		}
		ans = min(ans, cur)
		c += x
	}
	return int64(ans)

}

func main() {
	fmt.Println(minCost([]int{1, 2, 3}, 4))
	fmt.Println(minCost([]int{20, 1, 15}, 5))
}
