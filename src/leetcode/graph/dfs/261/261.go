package main

import "fmt"

/*
*
给定编号从 0 到 n - 1 的 n 个结点。给定一个整数 n 和一个 edges 列表，
其中 edges[i] = [ai, bi] 表示图中节点 ai 和 bi 之间存在一条无向边。

如果这些边能够形成一个合法有效的树结构，则返回 true ，否则返回 false 。
*/
func validTree(n int, edges [][]int) bool {

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := 0
	color := make([]int, n)
	var dfs func(int, int) bool

	dfs = func(x, fa int) bool {

		cnt++
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1

		for _, y := range g[x] {
			if y != fa {
				if !dfs(y, x) {
					return false
				}
			}
		}
		color[x] = 2
		return true
	}

	return dfs(0, -1) && cnt == n
}

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
}
