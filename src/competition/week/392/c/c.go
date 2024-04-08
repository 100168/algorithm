package main

import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {

	sort.Ints(nums)

	n := len(nums)

	minOps := int64(0)

	m := n / 2
	for i := 0; i < n; i++ {
		if i < m {
			if nums[i] > k {
				minOps += int64(nums[i] - k)
			}
		} else if i > m {
			if nums[i] < k {
				minOps += int64(k - nums[i])
			}
		}

	}
	minOps += int64(abs(k - nums[m]))
	return minOps
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
