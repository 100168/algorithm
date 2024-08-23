package main

import (
	"math"
	"slices"
)

/*
*
给你一个整数数组 nums 和两个整数 cost1 和 cost2 。你可以执行以下 任一 操作 任意 次：

从 nums 中选择下标 i 并且将 nums[i] 增加 1 ，开销为 cost1。
选择 nums 中两个 不同 下标 i 和 j ，并且将 nums[i] 和 nums[j] 都 增加 1 ，开销为 cost2 。
你的目标是使数组中所有元素都 相等 ，请你返回需要的 最小开销 之和。

由于答案可能会很大，请你将它对 109 + 7 取余 后返回。

示例 1：

输入：nums = [4,1], cost1 = 5, cost2 = 2

输出：15

解释：

执行以下操作可以使数组中所有元素相等：

将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,2] 。
将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,3] 。
将 nums[1] 增加 1 ，开销为 5 ，nums 变为 [4,4] 。
总开销为 15 。
*/
func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {

	mod := int(1e9 + 7)
	s := 0
	n := len(nums)
	maxVal := slices.Max(nums)
	minVal := slices.Min(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != maxVal {
			s += maxVal - nums[i]
		}
	}
	first := s * cost1
	second := math.MaxInt
	up := maxVal * 2

	for i := maxVal; i <= up; i++ {
		d := i - minVal
		cur := 0
		if 2*d-s <= 0 {
			cur = s/2*cost2 + (s%2)*cost1
		} else {
			cur = (s-d)*cost2 + cost1*(2*d-s)
		}
		s += n
		second = min(cur, second)
	}

	return min(first, second) % mod

}

func minCostToEqualizeArray2(nums []int, cost1 int, cost2 int) int {

	mod := int(1e9 + 7)
	maxV := slices.Max(nums)
	minV := slices.Min(nums)

	n := len(nums)

	s := 0
	for _, v := range nums {
		s += maxV - v

	}
	first := s * cost1

	up := maxV * 2

	ans := first

	for i := maxV; i <= up; i++ {

		d := i - minV
		cur := 0
		if 2*d <= s {
			cur = s/2*cost2 + (s%2)*cost1
		} else {
			cur = (s-d)*cost2 + (2*d-s)*cost1
		}
		ans = min(ans, cur)
		s += n

	}

	return ans % mod

}
func main() {
	println(minCostToEqualizeArray([]int{1, 14, 14, 15}, 2, 1))
	println(minCostToEqualizeArray([]int{7, 4, 8}, 7, 3))
	println(minCostToEqualizeArray2([]int{1, 14, 14, 15}, 2, 1))
	println(minCostToEqualizeArray2([]int{7, 4, 8}, 7, 3))

}
