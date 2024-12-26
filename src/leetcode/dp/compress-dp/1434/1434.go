package main

/*
*
共有 n 个人和 40 种不同的帽子，帽子编号从 1 到 40 。

给你一个整数列表的列表 hats ，其中 hats[i] 是第 i 个人所有喜欢帽子的列表。

请你给每个人安排一顶他喜欢的帽子，确保每个人戴的帽子跟别人都不一样，并返回方案数。

由于答案可能很大，请返回它对 10^9 + 7 取余后的结果。

示例 1：

输入：hats = [[3,4],[4,5],[5]]
输出：1
解释：给定条件下只有一种方法选择帽子。
第一个人选择帽子 3，第二个人选择帽子 4，最后一个人选择帽子 5。

tag：值域状压dp
*/
func numberWays(hats [][]int) int {

	n := len(hats)

	hMap := make(map[int][]int)

	for i, r := range hats {

		for _, v := range r {
			hMap[v] = append(hMap[v], i)
		}
	}

	mod := int(1e9 + 7)

	f := make([][]int, 41)

	for i := range f {
		f[i] = make([]int, 1<<n)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(i, mask int) int
	dfs = func(i, mask int) int {

		if mask == 1<<n-1 {
			return 1
		}
		if i < 1 {
			return 0
		}

		if f[i][mask] != -1 {
			return f[i][mask]
		}

		cur := 0

		for _, j := range hMap[i] {

			if mask>>j&1 == 0 {
				cur = (cur + dfs(i-1, mask|1<<j)) % mod
			}

		}
		cur = (cur + dfs(i-1, mask)) % mod

		f[i][mask] = cur
		return cur
	}
	return dfs(40, 0)

}
