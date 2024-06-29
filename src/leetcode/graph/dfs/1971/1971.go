package main

/*
*
有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）
。图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。

请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。

给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回 true，否则返回 false 。
*/
func validPath(n int, edges [][]int, source int, destination int) bool {

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(x, fa int) bool

	visited := make([]bool, n)
	dfs = func(x, fa int) bool {

		visited[x] = true
		if x == destination {
			return true
		}

		cur := false

		for _, y := range g[x] {
			if y == fa || visited[y] {
				continue
			}
			cur = cur || dfs(y, x)
		}

		return cur
	}

	return dfs(source, -1)
}
