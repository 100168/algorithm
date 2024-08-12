package main

import (
	"fmt"
	"sort"
)

/**
给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。

每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。

请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。



示例 1：

输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
输出：3
解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
- 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
- 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
- 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。

*/

func minOperations(nums1 []int, nums2 []int) int {

	m := len(nums1)
	n := len(nums2)

	if m*1 > n*6 || n*1 > m*6 {
		return -1
	}

	sort.Ints(nums2)
	sort.Ints(nums1)

	s1 := 0
	s2 := 0
	diff := 0
	for _, v := range nums1 {
		diff += v
		s1 += v
	}

	for _, v := range nums2 {
		diff -= v
		s2 += v
	}

	if diff == 0 {
		return 0
	}

	if diff < 0 {
		nums1, nums2 = nums2, nums1
		m, n = n, m
		diff = -diff
	}
	l, r := 0, m-1

	ops := 0
	for diff != 0 {
		if l == n || r >= 0 && nums1[r]-1 >= 6-nums2[l] {
			diff -= min(diff, nums1[r]-1)
			r--
			ops++
		} else {
			diff -= min(diff, 6-nums2[l])
			l++
			ops++
		}
	}
	return ops
}

func main() {
	//fmt.Println(minOperations([]int{5, 6, 4, 3, 1, 2}, []int{6, 3, 3, 1, 4, 5, 3, 4, 1, 3, 4}))
	//fmt.Println(minOperations([]int{1, 1, 1, 1, 1, 1, 1}, []int{6}))
	fmt.Println(minOperations([]int{6, 6}, []int{1}))
}
