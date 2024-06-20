package main

/*
*
你有一个包含 n 个节点的图。给定一个整数 n 和一个数组 edges ，
其中 edges[i] = [ai, bi] 表示图中 ai 和 bi 之间有一条边。
返回 图中已连接分量的数目 。
*/
func countComponents(n int, edges [][]int) int {

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	visited := make([]bool, n)

	var dfs func(int)

	dfs = func(x int) {

		visited[x] = true
		for _, y := range g[x] {
			if !visited[y] {
				dfs(y)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			ans++
		}
	}
	return ans
}
