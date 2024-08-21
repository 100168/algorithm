package main

import "fmt"

/**
给定一个整数数组 nums 和一个整数 k，你可以对数组执行以下操作任意次数：

选择数组中的两个 相邻 元素，例如 x 和 y，使得 x * y <= k ，
并用一个值为 x * y 的 单个元素 替换它们（例如，在一次操作中，数组 [1, 2, 2, 3]，
其中 k = 5 可以变为 [1, 4, 3] 或 [2, 2, 3]，但不能变为 [1, 2, 6]）。
返回 经过任意次数的操作后， nums 的 最小 可能长度。

示例 1：

输入：nums = [2,3,3,7,3,5], k = 20
输出：3
解释：我们执行以下操作：
1. [2,3,3,7,3,5] -> [6,3,7,3,5]
2. [6,3,7,3,5] -> [18,7,3,5]
3. [18,7,3,5] -> [18,7,15]
可以证明，在执行给定操作后，最小可能长度为3.
*/

func minArrayLength(nums []int, k int) int {

	ans := 1
	pre := nums[0]
	for _, v := range nums[1:] {
		if v == 0 {
			return 1
		}
		if pre*v > k {
			pre = v
			ans++
			continue
		}
		pre *= v
	}
	return ans
}

func main() {
	fmt.Println(minArrayLength([]int{1, 5}, 4))
	fmt.Println(minArrayLength([]int{7}, 1))
}
