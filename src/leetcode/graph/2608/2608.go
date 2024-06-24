package main

import (
	"fmt"
	"math"
)

/*
现有一个含 n 个顶点的 双向 图，每个顶点按从 0 到 n - 1 标记。
图中的边由二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和 vi 之间存在一条边
。每对顶点最多通过一条边连接，并且不存在与自身相连的顶点。

返回图中 最短 环的长度。如果不存在环，则返回 -1 。

环 是指以同一节点开始和结束，并且路径中的每条边仅使用一次。
*/
func findShortestCycle(n int, edges [][]int) int {

	g := make([][]int, n)
	for _, v := range edges {

		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dis := make([]int, n)

	ans := math.MaxInt

	type pair struct{ x, fa int }
	for i := 0; i < n; i++ {
		for j := range dis {
			dis[j] = -1
		}
		dis[i] = 0
		queue := []pair{{i, -1}}
		for len(queue) > 0 {
			x, fa := queue[0].x, queue[0].fa
			queue = queue[1:]
			for _, y := range g[x] {
				if dis[y] == -1 {
					dis[y] = dis[x] + 1
					queue = append(queue, pair{y, x})
				} else if fa != y {
					ans = min(ans, dis[x]+dis[y]+1)
				}
			}

		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans

}

func findShortestCycle2(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	ans := math.MaxInt
	dis := make([]int, n)                // dis[i] 表示从 start 到 i 的最短路长度
	for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
		for j := range dis {
			dis[j] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		q := []pair{{start, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, fa := p.x, p.fa
			for _, y := range g[x] {
				if dis[y] < 0 { // 第一次遇到
					dis[y] = dis[x] + 1
					q = append(q, pair{y, x})
				} else if y != fa { // 第二次遇到
					ans = min(ans, dis[x]+dis[y]+1)
				}
			}
		}
	}
	if ans == math.MaxInt { // 无环图
		return -1
	}
	return ans
}

func main() {
	fmt.Println(findShortestCycle(8, [][]int{{1, 3}, {3, 5}, {5, 7}, {7, 1}, {0, 2}, {2, 4}, {4, 0}, {6, 0}, {6, 1}}))
}
