package main

import (
	"fmt"
	"math"
)

/*
*
给定一个正整数 n，表示一个 连通无向图 中的节点数，该图 只包含一个 环。节点编号为 0 ~ n - 1(含)。

你还得到了一个二维整数数组 edges，其中 edges[i] = [node1i, node2i] 表示有一条 双向 边连接图中的 node1i 和 node2i。

两个节点 a 和 b 之间的距离定义为从 a 到 b 所需的 最小边数。

返回一个长度为 n 的整数数组 answer，其中 answer[i] 是第 i 个节点与环中任何节点之间的最小距离。
*/
func distanceToCycle(n int, edges [][]int) []int {

	g := make([][]int, n)

	inDegree := make([]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		inDegree[x]++
		inDegree[y]++
	}
	var queue []int

	for i, v := range inDegree {
		if v == 1 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, y := range g[x] {
			inDegree[y]--
			if inDegree[y] == 1 {
				queue = append(queue, y)
			}
		}
	}
	ans := make([]int, n)

	visited := make([]bool, n)

	var st []int
	for i, v := range inDegree {
		if v == 2 {
			visited[i] = true
			st = append(st, i)
		}
	}

	for i := range ans {
		ans[i] = math.MaxInt
		if visited[i] {
			ans[i] = 0
		}
	}
	for len(st) > 0 {
		cur := st[0]
		st = st[1:]
		for _, v := range g[cur] {
			if !visited[v] {
				visited[v] = true
				ans[v] = ans[cur] + 1
				st = append(st, v)
			}
		}
	}

	return ans

}

func distanceToCycle2(n int, edges [][]int) []int {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v) // 建图
		deg[v]++
		deg[w]++
	}

	// 拓扑排序，剪掉所有树枝
	var q []int
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 1 {
				q = append(q, w)
			}
		}
	}

	// 从基环出发，求所有树枝上的点的深度
	ans := make([]int, n)
	var f func(int, int)
	f = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa && deg[w] < 2 {
				ans[w] = ans[v] + 1
				f(w, v)
			}
		}
	}
	for root, d := range deg {
		if d > 1 {
			f(root, -1)
		}
	}
	return ans
}

func main() {

	fmt.Println(distanceToCycle(7, [][]int{{1, 2}, {2, 4}, {4, 3}, {3, 1}, {0, 1}, {5, 2}, {6, 5}}))
}
