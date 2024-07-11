package main

import "math"

/*
*
给你一个无向图，整数 n 表示图中节点的数目，edges 数组表示图中的边，其中 edges[i] = [ui, vi] ，表示 ui 和 vi 之间有一条无向边。

一个 连通三元组 指的是 三个 节点组成的集合且这三个点之间 两两 有边。

连通三元组的度数 是所有满足此条件的边的数目：一个顶点在这个三元组内，而另一个顶点不在这个三元组内。

请你返回所有连通三元组中度数的 最小值 ，如果图中没有连通三元组，那么返回 -1
*/
func minTrioDegree(n int, edges [][]int) int {

	s := make([]int, n+1)
	g := make([][]int, n+1)
	for i := range g {
		g[i] = make([]int, n+1)
	}
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x][y] = 1
		g[y][x] = 1
		s[x]++
		s[y]++
	}
	ans := math.MaxInt
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if g[i][j] == 0 {
				continue
			}
			if i == j {
				continue
			}
			for k := 1; k <= n; k++ {
				if g[i][k] == 0 || g[j][k] == 0 {
					continue
				}
				ans = min(ans, s[i]+s[j]+s[k]-6)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
