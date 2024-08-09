package main

/**

给你一个整数 n 和一个二维整数数组 queries。

有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。


*/

func shortestDistanceAfterQueries(n int, queries [][]int) []int {

	ans := make([]int, len(queries))
	g := make(map[int][]int, n)
	for i := 0; i < n-1; i++ {
		g[i] = append(g[i], i+1)
	}
next:
	for i, q := range queries {
		x, y := q[0], q[1]
		g[x] = append(g[x], y)
		st := []int{0}
		d := 0

		visited := make([]bool, n)
		visited[0] = true
		for len(st) > 0 {
			temp := st
			st = nil
			for _, c := range temp {
				if c == n-1 {
					ans[i] = d
					continue next
				}
				for _, next := range g[c] {
					if !visited[next] {
						st = append(st, next)
					}
					visited[next] = true
				}
			}
			d++
		}

	}
	return ans
}
