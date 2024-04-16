package main

import "sort"

// 给你一个下标从 0 开始的正整数数组 nums 。请你找出并统计满足下述条件的三元组 (i, j, k) 的数目：
//
// 0 <= i < j < k < nums.length
// nums[i]、nums[j] 和 nums[k] 两两不同 。
// 换句话说：nums[i] != nums[j]、nums[i] != nums[k] 且 nums[j] != nums[k] 。
// 返回满足上述条件三元组的数目。
func unequalTriplets(nums []int) (ans int) {
	//先排序
	sort.Ints(nums)
	//再分组
	start, n := 0, len(nums)
	for i, x := range nums[:n-1] {
		if x != nums[i+1] {
			ans += start * (i - start + 1) * (n - 1 - i)
			start = i + 1
		}
	}
	return
}
