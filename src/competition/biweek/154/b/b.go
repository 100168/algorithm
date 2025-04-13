package main

import (
	"fmt"
	"math/bits"
)

/**
给你一个长度为 n 的整数数组 nums，其中 nums 是范围 [1, n] 内所有数的排列。

XOR 三元组 定义为三个元素的异或值 nums[i] XOR nums[j] XOR nums[k]，其中 i <= j <= k。

返回所有可能三元组 (i, j, k) 中不同的 XOR 值的数量。

排列 是一个集合中所有元素的重新排列。



示例 1：

输入： nums = [1,2]

输出： 2

解释：

所有可能的 XOR 三元组值为：

(0, 0, 0) → 1 XOR 1 XOR 1 = 1
(0, 0, 1) → 1 XOR 1 XOR 2 = 2
(0, 1, 1) → 1 XOR 2 XOR 2 = 1
(1, 1, 1) → 2 XOR 2 XOR 2 = 2
不同的 XOR 值为 {1, 2}，因此输出为 2。
*/

func uniqueXorTriplets(nums []int) int {

	n := len(nums)
	if n < 3 {
		return len(nums)
	}

	l := bits.Len(uint(n))

	return 1 << l
}

func main() {
	fmt.Println(uniqueXorTriplets([]int{1, 2}))
}
