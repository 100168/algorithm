package main

/*
*给你一个 n 个点的 简单有向图 （没有重复边的有向图），节点编号为 0 到 n - 1 。如果这些边是双向边，那么这个图形成一棵 树 。

给你一个整数 n 和一个 二维 整数数组 edges ，其中 edges[i] = [ui, vi] 表示从节点 ui 到节点 vi 有一条 有向边 。

边反转 指的是将一条边的方向反转，也就是说一条从节点 ui 到节点 vi 的边会变为一条从节点 vi 到节点 ui 的边。

对于范围 [0, n - 1] 中的每一个节点 i ，你的任务是分别 独立 计算 最少 需要多少次 边反转 ，
从节点 i 出发经过 一系列有向边 ，可以到达所有的节点。

请你返回一个长度为 n 的整数数组 answer ，其中 answer[i]表示从节点 i 出发，可以到达所有节点的 最少边反转 次数。
*/
func minEdgeReversals(n int, edges [][]int) []int {

	g := make([][]int, n)

	type pair struct{ x, y int }

	dirs := make(map[pair]bool)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		dirs[pair{x, y}] = true

	}

	ans := make([]int, n)
	var dfs func(int, int)

	dfs = func(x, fa int) {

		for _, y := range g[x] {

			if y == fa {
				continue
			}
			if dirs[pair{y, x}] {
				ans[0]++
			}
			dfs(y, x)
		}
	}

	dfs(0, -1)

	var reRoot func(int, int)

	reRoot = func(x, fa int) {

		for _, y := range g[x] {

			if y == fa {
				continue
			}
			cur := ans[x]
			if dirs[pair{x, y}] {
				cur++
			} else {
				cur--
			}
			ans[y] = cur
			reRoot(y, x)
		}
	}
	reRoot(0, -1)
	return ans
}
