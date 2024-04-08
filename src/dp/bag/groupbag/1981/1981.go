package main

import (
	"math"
	"slices"
	"sort"
)

//给你一个大小为 m x n 的整数矩阵 mat 和一个整数 target 。
//
// 从矩阵的 每一行 中选择一个整数，你的目标是 最小化 所有选中元素之 和 与目标值 target 的 绝对差 。
//
// 返回 最小的绝对差 。
//
// a 和 b 两数字的 绝对差 是 a - b 的绝对值。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[1,2,3],[4,5,6],[7,8,9]], target = 13
//输出：0
//解释：一种可能的最优选择方案是：
//- 第一行选出 1
//- 第二行选出 5
//- 第三行选出 7
//所选元素的和是 13 ，等于目标值，所以绝对差是 0 。
//
//
// 示例 2：
//
//
//
//
//输入：mat = [[1],[2],[3]], target = 100
//输出：94
//解释：唯一一种选择方案是：
//- 第一行选出 1
//- 第二行选出 2
//- 第三行选出 3
//所选元素的和是 6 ，绝对差是 94 。
//
//
// 示例 3：
//
//
//
//
//输入：mat = [[1,2,9,8,7]], target = 6
//输出：1
//解释：最优的选择方案是选出第一行的 7 。
//绝对差是 1 。
//
//
//
//
// 提示：
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 70
// 1 <= mat[i][j] <= 70
// 1 <= target <= 800
//
//
// Related Topics 数组 动态规划 矩阵 👍 66 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func minimizeTheDifference(mat [][]int, target int) int {

	m := len(mat)
	cache := make([][]int, m)
	for i := range mat {
		sort.Ints(mat[i])
	}
	for i := range cache {
		cache[i] = make([]int, 5000)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= m {
			return abs(j - target)
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		cur := math.MaxInt
		for _, v := range mat[i] {
			cur = min(cur, dfs(i+1, v+j))
			if v+j > target {
				break
			}
		}
		cache[i][j] = cur
		return cur
	}
	return dfs(0, 0)
}

func abs(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

func dp(mat [][]int, target int) int {
	dp := make([]bool, min(len(mat)*70, target*2)+1) // 需要枚举的重量不会超过 target*2
	dp[0] = true
	minSum, maxSum := 0, 0
	for _, row := range mat {
		mi, mx := slices.Min(row), slices.Max(row)
		minSum += mi                      // 求 minSum 是为了防止 target 过小导致 dp 没有记录到
		maxSum = min(maxSum+mx, target*2) // 前 i 组的最大重量，优化枚举时 j 的初始值
		for j := maxSum; j >= 0; j-- {
			dp[j] = false
			for _, v := range row {
				if v <= j && dp[j-v] {
					dp[j] = true
					break
				}
			}
		}
	}
	ans := abs(minSum - target)
	for i, ok := range dp {
		if ok {
			ans = min(ans, abs(i-target))
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
