package main

import "math"

/**
给你一个大小为 4 的整数数组 a 和一个大小 至少为 4 的整数数组 b。

你需要从数组 b 中选择四个下标 i0, i1, i2, 和 i3，并满足 i0 < i1 < i2 < i3。
你的得分将是 a[0] * b[i0] + a[1] * b[i1] + a[2] * b[i2] + a[3] * b[i3] 的值。

返回你能够获得的 最大 得分。



示例 1：

输入： a = [3,2,5,6], b = [2,-6,4,-5,-3,2,-7]

输出： 26

解释：
选择下标 0, 1, 2 和 5。得分为 3 * 2 + 2 * (-6) + 5 * 4 + 6 * 2 = 26。

示例 2：

输入： a = [-1,4,5,-2], b = [-5,-1,-3,-2,-4]

输出： -1

解释：
选择下标 0, 1, 3 和 4。得分为 (-1) * (-5) + 4 * (-1) + 5 * (-2) + (-2) * (-4) = -1。
*/

func maxScore(a []int, b []int) int64 {

	n := len(b)

	f := make([][]int, n)

	for i := range f {
		f[i] = make([]int, 4)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i int, rest int) int {

		if rest < 0 {
			return 0
		}
		if i < 0 {
			return math.MinInt / 2
		}

		if f[i][rest] != -1 {
			return f[i][rest]
		}

		c := max(dfs(i-1, rest), dfs(i-1, rest-1)+a[rest]*b[i])

		f[i][rest] = c
		return c

	}

	return int64(dfs(n-1, 3))
}
