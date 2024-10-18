package main

import (
	"slices"
)

/*
*
给你一个长度为 n的整数数组nums和一个整数数组queries。

gcdPairs表示数组 nums中所有满足 0 <= i < j < n的数对 (nums[i], nums[j]) 的
最大公约数升序排列构成的数组。

对于每个查询queries[i]，你需要找到gcdPairs中下标为queries[i]的元素。

请你返回一个整数数组answer，其中answer[i]是gcdPairs[queries[i]]的值。

gcd(a, b)表示 a和 b的 最大公约数。

示例 1：

输入：nums = [2,3,4], queries = [0,2,2]

输出：[1,2,2]

解释：

gcdPairs = [gcd(nums[0], nums[1]), gcd(nums[0], nums[2]), gcd(nums[1], nums[2])] = [1, 2, 1].

升序排序后得到gcdPairs = [1, 1, 2]。

所以答案为[gcdPairs[queries[0]], gcdPairs[queries[1]], gcdPairs[queries[2]]] = [1, 2, 2]。
*/
func gcdValues(nums []int, queries []int64) []int {
	cntMap := make(map[int]int)

	for _, v := range nums {
		cntMap[v]++
	}
	maxV := slices.Max(nums)

	cnt := make([]int, maxV+1)
	sum := make([]int, maxV+2)

	//2,4,6,8

}
