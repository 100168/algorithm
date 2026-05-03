package main

/*
给你一个整数数组 nums。

在一次操作中，你需要移除当前数组的 前三个元素。如果剩余元素少于三个，则移除 所有 剩余元素。

重复此操作，直到数组为空或不包含任何重复元素为止。

返回一个整数，表示所需的操作次数。

示例 1:

输入: nums = [3,8,3,6,5,8]

输出: 1

解释:

在第一次操作中，我们移除前三个元素。剩余的元素 [6, 5, 8] 互不相同，因此停止。仅需要一次操作。
*/

func minOperations(nums []int) int {
	n := len(nums)
	seen := make(map[int]bool)
	for i := n - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return i/3 + 1
		}
		seen[nums[i]] = true
	}
	return 0
}

func main() {
	println(minOperations([]int{3, 8, 3, 3, 3, 8}))
}
