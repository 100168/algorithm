package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个整数数组 nums 和两个整数 k 和 numOperations 。

你必须对 nums 执行 操作  numOperations 次。每次操作中，你可以：

选择一个下标 i ，它在之前的操作中 没有 被选择过。
将 nums[i] 增加范围 [-k, k] 中的一个整数。
在执行完所有操作以后，请你返回 nums 中出现 频率最高 元素的出现次数。

一个元素 x 的 频率 指的是它在数组中出现的次数。

示例 1：

输入：nums = [1,4,5], k = 1, numOperations = 2

输出：2

解释：

通过以下操作得到最高频率 2 ：

将 nums[1] 增加 0 ，nums 变为 [1, 4, 5] 。
将 nums[2] 增加 -1 ，nums 变为 [1, 4, 4] 。
示例 2：

输入：nums = [5,11,20,20], k = 5, numOperations = 1

输出：2

解释：

通过以下操作得到最高频率 2 ：

将 nums[1] 增加 0 。

提示：

1 <= nums.length <= 105
1 <= nums[i] <= 105
0 <= k <= 105
0 <= numOperations <= nums.length
*/
func maxFrequency(nums []int, k int, numOperations int) int {

	sort.Ints(nums)
	cnt := make(map[int]int)

	for _, v := range nums {
		cnt[v]++
	}
	find := func(t int) int {

		l, r := 0, len(nums)-1

		for l <= r {

			m := (l + r) / 2

			if nums[m] >= t {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return l
	}
	ans := 1

	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		l := find(i - k)
		r := find(i + k + 1)
		ans = max(ans, min(r-l-cnt[i], numOperations)+cnt[i])
	}
	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4, 5}, 2, 4))
}
