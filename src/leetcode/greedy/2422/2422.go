package main

/**
给定一个由 正整数 组成的数组 nums。

可以对阵列执行如下操作，次数不限:

选择任意两个 相邻 的元素并用它们的 和 替换 它们。
例如，如果 nums = [1,2,3,1]，则可以应用一个操作使其变为 [1,5,1]。
返回将数组转换为 回文序列 所需的 最小 操作数。



示例 1:

输入: nums = [4,3,2,1,2,3,1]
输出: 2
解释: 我们可以通过以下 2 个操作将数组转换为回文:
- 在数组的第 4 和第 5 个元素上应用该操作，nums 将等于 [4,3,2,3,3,1].
- 在数组的第 5 和第 6 个元素上应用该操作，nums 将等于 [4,3,2,3,4].
数组 [4,3,2,3,4] 是一个回文序列。
可以证明，2 是所需的最小操作数。

贪心

*/

func minimumOperations(nums []int) int {

	n := len(nums)

	if n == 1 {
		return 0
	}

	l, r := 0, n-1

	ans := 0

	for l < r {
		if nums[l] < nums[r] {
			ans++
			nums[l+1] += nums[l]
			l++
		} else if nums[r] < nums[l] {
			ans++
			nums[r-1] += nums[r]
			r--
		} else {
			l++
			r--
		}
	}
	return ans
}
