package main

import "slices"

/*
*给你一个下标从 0 开始、长度为 n 的整数数组 nums ，和一个整数 k 。

你可以执行下述 递增 运算 任意 次（可以是 0 次）：

从范围 [0, n - 1] 中选择一个下标 i ，并将 nums[i] 的值加 1 。
如果数组中任何长度 大于或等于 3 的子数组，其 最大 元素都大于或等于 k ，则认为数组是一个 美丽数组 。

dp[i][j]
以整数形式返回使数组变为 美丽数组 需要执行的 最小 递增运算数。

子数组是数组中的一个连续 非空 元素序列。

示例 1：

输入：nums = [2,3,0,0,2], k = 4
输出：3
解释：可以执行下述递增运算，使 nums 变为美丽数组：
选择下标 i = 1 ，并且将 nums[1] 的值加 1 -> [2,4,0,0,2] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,3] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,4] 。
长度大于或等于 3 的子数组为 [2,4,0], [4,0,0], [0,0,4], [2,4,0,0], [4,0,0,4], [2,4,0,0,4] 。
在所有子数组中，最大元素都等于 k = 4 ，所以 nums 现在是美丽数组。
可以证明无法用少于 3 次递增运算使 nums 变为美丽数组。
因此，答案为 3 。
*/
func minIncrementOperations(nums []int, k int) int64 {

	n := len(nums)

	memo := make([][3]int64, n)
	for i := range memo {
		memo[i] = [3]int64{-1, -1, -1}
	}

	var dfs func(int, int) int64

	dfs = func(i int, rest int) int64 {
		if i < 0 {
			return 0
		}

		cur := int64(max(k-nums[i], 0))

		if memo[i][rest] != -1 {
			return memo[i][rest]
		}
		if rest == 0 {
			cur += dfs(i-1, 2)
		} else {
			cur = min(dfs(i-1, 2)+cur, dfs(i-1, rest-1))
		}
		memo[i][rest] = cur
		return cur
	}
	return dfs(n-1, 2)
}
func minIncrementOperations2(nums []int, k int) int64 {
	f := make([]int64, 3)
	for i := range nums {
		cur := int64(max(k-nums[i], 0))
		f[0], f[1], f[2] = f[2]+cur, min(f[2]+cur, f[0]), min(f[2]+cur, f[1])
	}
	return slices.Min(f)
}

func main() {
	println(minIncrementOperations([]int{1, 1, 2, 1}, 3))
	println(minIncrementOperations2([]int{1, 1, 2, 1}, 3))
}
