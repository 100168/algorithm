package main

import (
	"fmt"
	"math"
)

/*
*
给你一个整数数组 nums 和一个整数 k 。

将数组拆分成一些非空子数组。拆分的 代价 是每个子数组中的 重要性 之和。

令 trimmed(subarray) 作为子数组的一个特征，其中所有仅出现一次的数字将会被移除。

例如，trimmed([3,1,2,4,3,4]) = [3,4,3,4] 。
子数组的 重要性 定义为 k + trimmed(subarray).length 。

例如，如果一个子数组是 [1,2,3,3,3,4,4] ，trimmed([1,2,3,3,3,4,4]) = [3,3,3,4,4] 。这个子数组的重要性就是 k + 5 。
找出并返回拆分 nums 的所有可行方案中的最小代价。

子数组 是数组的一个连续 非空 元素序列。

输入：nums = [1,2,1,2,1,3,3], k = 2
输出：8
解释：将 nums 拆分成两个子数组：[1,2], [1,2,1,3,3]
[1,2] 的重要性是 2 + (0) = 2 。
[1,2,1,3,3] 的重要性是 2 + (2 + 2) = 6 。
拆分的代价是 2 + 6 = 8 ，可以证明这是所有可行的拆分方案中的最小代价。

输入：nums = [1,2,1,2,1], k = 5
输出：10
解释：将 nums 拆分成一个子数组：[1,2,1,2,1].
[1,2,1,2,1] 的重要性是 5 + (3 + 2) = 10 。
*/
func minCost(nums []int, k int) int {

	n := len(nums)

	f := make([]int, n)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if f[i] != -1 {
			return f[i]
		}
		cur := math.MaxInt
		cnt := 0
		cntMap := make(map[int]int)

		for j := i; j >= 0; j-- {

			if cntMap[nums[j]]++; cntMap[nums[j]] > 1 {
				if cntMap[nums[j]] == 2 {
					cnt++
				}
				cnt++
			}
			cur = min(cur, dfs(j-1)+k+cnt)
		}
		f[i] = cur
		return cur

	}
	return dfs(n - 1)
}

func main() {
	fmt.Println(minCost([]int{1, 2, 1, 2, 1, 3, 3}, 2))
}
