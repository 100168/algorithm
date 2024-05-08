package main

import (
	"slices"
)

/*
*

给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。

你可以对数组执行 至多 k 次操作：

从数组中选择一个下标 i ，将 nums[i] 增加 或者 减少 1 。
最终数组的频率分数定义为数组中众数的 频率 。

请你返回你可以得到的 最大 频率分数。

众数指的是数组中出现次数最多的数。一个元素的频率指的是数组中这个元素的出现次数。

1,2,3,4,5    3
*/
func maxFrequencyScore(nums []int, K int64) (ans int) {
	k := int(K)
	slices.Sort(nums)

	n := len(nums)
	sum := make([]int, n+1)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}

	//中位数贪心
	// 把 nums[l] 到 nums[r] 都变成 nums[i]
	distanceSum := func(l, i, r int) int {
		left := nums[i]*(i-l) - (sum[i] - sum[l])
		right := sum[r+1] - sum[i+1] - nums[i]*(r-i)
		return left + right
	}

	left := 0
	for i := range nums {
		for distanceSum(left, (left+i)/2, i) > k {
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}

func maxFrequencyScore2(nums []int, k int64) int {
	slices.Sort(nums)
	ans, left := 0, 0
	s := int64(0) // 窗口元素与窗口中位数的差之和
	for right, x := range nums {
		s += int64(x - nums[(left+right)/2])
		for s > k {
			s += int64(nums[left] - nums[(left+right+1)/2])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
