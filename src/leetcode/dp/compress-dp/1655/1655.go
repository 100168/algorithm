package main

import "sort"

/*
*

给你一个长度为 n 的整数数组 nums ，这个数组中至多有 50 个不同的值。同时你有 m 个顾客的订单 quantity
，其中，整数 quantity[i] 是第 i 位顾客订单的数目。请你判断是否能将 nums 中的整数分配给这些顾客，且满足：

第 i 位顾客 恰好 有 quantity[i] 个整数。
第 i 位顾客拿到的整数都是 相同的 。
每位顾客都满足上述两个要求。
如果你可以分配 nums 中的整数满足上面的要求，那么请返回 true ，否则返回 false 。
*/
func canDistribute(nums []int, quantity []int) bool {

	//本质上跟分糖果是一样的
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}

	m := len(quantity)
	newNums := make([]int, 0)
	for _, v := range numsMap {
		newNums = append(newNums, v)
	}

	sum := make([]int, 1<<m)
	for i, v := range quantity {
		for j, k := 0, 1<<i; j < k; j++ {
			sum[j|k] = sum[j] + v
		}
	}

	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i] > newNums[j]
	})

	newNums = newNums[:min(len(newNums), m)]

	n := len(newNums)

	dp := make([]bool, 1<<m)

	for i := range dp {
		dp[i] = false
	}

	dp[0] = true

	for i := 0; i < min(m, n); i++ {
		for s := 1<<m - 1; s > 0; s-- {
			for sb := s; sb > 0; sb = (sb - 1) & s {
				dp[s] = dp[s] || (newNums[i] >= sum[sb] && dp[s^sb])
			}
		}
	}

	return dp[1<<m-1]

}

func main() {
	println(canDistribute([]int{357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357, 357}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
