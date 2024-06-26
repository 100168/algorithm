package main

/*
*
给你一个整数 n 。现有一个包含 n 个顶点的 无向 图，顶点按从 0 到 n - 1 编号。
给你一个二维整数数组 edges 其中 edges[i] = [ai, bi] 表示顶点 ai 和 bi 之间存在一条 无向 边。

返回图中 完全连通分量 的数量。

如果在子图中任意两个顶点之间都存在路径，并且子图中没有任何一个顶点与子图外部的顶点共享边，则称其为 连通分量 。

如果连通分量中每对节点之间都存在一条边，则称其为 完全连通分量 。
*/
func countCompleteComponents(n int, edges [][]int) int {

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	visited := make([]bool, n)
	cnt := 0
	v := 0
	var dfs func(int)
	dfs = func(x int) {
		visited[x] = true
		v++
		cnt += len(g[x])

		for _, y := range g[x] {
			if !visited[y] {
				dfs(y)
			}
		}
	}
	ans := 0

	for i := 0; i < n; i++ {
		cnt, v = 0, 0
		if !visited[i] {
			dfs(i)
			if cnt == v*(v-1) {
				ans++
			}
		}
	}
	return ans
}
