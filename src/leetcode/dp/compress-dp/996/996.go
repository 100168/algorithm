package main

import (
	"fmt"
	"math"
)

/*
*

给定一个非负整数数组 A，如果该数组每对相邻元素之和是一个完全平方数，则称这一数组为正方形数组。

返回 A 的正方形排列的数目。两个排列 A1 和 A2 不同的充要条件是存在某个索引 i，使得 A1[i] != A2[i]。
*/
func numSquarefulPerms(nums []int) int {

	n := len(nums)
	memo := make([]map[int]int, 1<<n)
	for i := range memo {
		memo[i] = make(map[int]int)
	}

	var dfs func(int, int) int

	dfs = func(mask, pre int) int {

		if mask == 1<<n-1 {
			return 1
		}
		if _, ok := memo[mask][pre]; ok {
			return memo[mask][pre]
		}

		cur := 0
		check := make(map[int]bool)
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 && (isSqrt(pre+nums[j]) || pre == -1) && !check[nums[j]] {
				check[nums[j]] = true
				cur += dfs(mask|1<<j, nums[j])
			}
		}
		memo[mask][pre] = cur
		return cur
	}

	return dfs(0, -1)
}

func isSqrt(a int) bool {
	sqrt := int(math.Sqrt(float64(a)))
	return sqrt*sqrt == a
}

func main() {
	nums := []int{3}

	fmt.Println(numSquarefulPerms(nums))
}
