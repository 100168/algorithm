package main

import (
	"fmt"
)

/**
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 k 。每一次操作中，你可以选择一个数并将它乘 2 。

你最多可以进行 k 次操作，请你返回 nums[0] | nums[1] | ... | nums[n - 1] 的最大值。

a | b 表示两个整数 a 和 b 的 按位或 运算。


示例 1：

输入：nums = [12,9], k = 1
输出：30
解释：如果我们对下标为 1 的元素进行操作，新的数组为 [12,18] 。此时得到最优答案为 12 和 18 的按位或运算的结果，也就是 30 。
示例 2：

输入：nums = [8,1,2], k = 2
输出：35
解释：如果我们对下标 0 处的元素进行操作，得到新数组 [32,1,2] 。此时得到最优答案为 32|1|2 = 35 。


提示：

1 <= nums.length <= 105
1 <= nums[i] <= 109
1 <= k <= 15
*/

func maximumOr(nums []int, k int) int64 {

	ans := 0

	cnt := make([]int, 32)
	cur := 0
	for _, v := range nums {
		for j := 0; j < 32; j++ {
			if 1<<j&v != 0 {
				cnt[j]++
			}
		}
		cur |= v
	}

	for _, v := range nums {

		temp := cur
		for j := 0; j < 32; j++ {
			if 1<<j&v != 0 && cnt[j] == 1 {
				temp ^= 1 << j
			}
		}
		temp |= v << k
		ans = max(ans, temp)

	}

	return int64(ans)

}

func maximumOr2(nums []int, k int) int64 {

	n := len(nums)
	//前后缀分解
	suffix := make([]int, n)
	suffix[n-1] = 0
	for i := n - 2; i >= 0; i-- {

		suffix[i] = suffix[i+1] | nums[i+1]

	}

	pre := 0
	ans := 0

	for i, v := range nums {

		ans = max(ans, pre|suffix[i]|v<<k)
		pre |= v
	}

	return int64(ans)

}

func main() {
	fmt.Println(maximumOr([]int{12, 9}, 1))
	fmt.Println(maximumOr([]int{4, 100, 76, 37, 99, 79, 39}, 4))
}
