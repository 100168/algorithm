package main

import (
	"fmt"
)

/*
*
给你一个下标从 1 开始、由 n 个整数组成的数组。你需要从 nums 选择一个 完全集，其中每对元素下标的乘积都是一个
完全平方数
，例如选择 ai 和 aj ，i * j 一定是完全平方数。

返回 完全子集 所能取到的 最大元素和 。

示例 1：

输入：nums = [8,7,3,5,7,2,4,9]
输出：16
解释：我们选择了下标 2 和 8 的元素，并且 2 * 8 是一个完全平方数。
示例 2：

输入：nums = [8,10,3,8,1,13,7,9,4]
输出：20
解释：我们选择了下标 1，4 和 9 的元素。1 * 4，1 * 9，4 * 9 都是完全平方数。
*/
func maximumSum(nums []int) int64 {

	n := len(nums)
	primeMap := make(map[int]int)

	for i := 1; i <= n; i++ {
		cur := i
		core := 1

		for j := 2; j*j <= cur; j++ {
			cnt := 0
			if cur%j == 0 {
				for cur%j == 0 {
					cnt++
					cur /= j
				}
				if cnt%2 == 1 {
					core *= j
				}
			}
		}
		if cur > 1 {
			core *= cur
		}
		primeMap[core] += nums[i-1]
	}
	ans := 0
	for _, v := range primeMap {
		ans = max(ans, v)
	}
	return int64(ans)

}

func main() {
	//fmt.Println(maximumSum([]int{8, 7, 3, 5, 7, 2, 4, 9}))
	//fmt.Println(maximumSum([]int{19, 6, 30, 23, 25, 45, 15, 2, 3, 46}))
	fmt.Println(maximumSum([]int{23, 27, 42, 3, 33, 36, 43, 32, 27, 48, 40, 22, 5, 36, 48}))
}
