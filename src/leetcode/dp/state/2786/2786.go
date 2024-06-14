package main

import (
	"fmt"
	"math"
)

/*
*
给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。

你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：

如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
对于你访问的位置 i ，你可以获得分数 nums[i] 。
如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。
请你返回你能得到的 最大 得分之和。

注意 ，你一开始的分数为 nums[0] 。

输入：nums = [2,3,6,1,9,2], x = 5
输出：13
解释：我们可以按顺序访问数组中的位置：0 -> 2 -> 3 -> 4 。
对应位置的值为 2 ，6 ，1 和 9 。因为 6 和 1 的奇偶性不同，所以下标从 2 -> 3 让你失去 x = 5 分。
总得分为：2 + 6 + 1 + 9 - 5 = 13 。
*/
func maxScore(nums []int, x int) int64 {

	odd := math.MinInt / 2
	even := math.MinInt / 2

	if nums[0]%2 == 0 {
		even = nums[0]
	} else {
		odd = nums[0]
	}

	for _, v := range nums[1:] {
		if v%2 == 0 {
			even = max(even, odd-x) + v
		} else {
			odd = max(odd, even-x) + v
		}

	}
	return int64(max(odd, even))
}

func main() {
	fmt.Println(maxScore([]int{2, 3, 6, 1, 9, 2}, 5))
}
