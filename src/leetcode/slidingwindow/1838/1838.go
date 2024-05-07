package main

import "sort"

/*
*
元素的 频数 是该元素在一个数组中出现的次数。

给你一个整数数组 nums 和一个整数 k 。在一步操作中，你可以选择 nums 的一个下标，并将该下标对应元素的值增加 1 。

执行最多 k 次操作后，返回数组中最高频元素的 最大可能频数 。
*/
func maxFrequency(nums []int, k int) int {

	n := len(nums)

	w := 0

	ans := 0
	l := 0
	sort.Ints(nums)
	pre := 0
	for r := 0; r < n; r++ {

		w += (nums[r] - pre) * (r - l)

		for w > k {
			w -= nums[r] - nums[l]
			l++
		}
		pre = nums[r]
		ans = max(ans, r-l+1)

	}
	return ans
}
