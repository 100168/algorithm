package main

import (
	"fmt"
	"math"
)

/*
*
给定一个整数 n，即有向图中的节点数，其中节点标记为 0 到 n - 1。图中的每条边为红色或者蓝色，并且可能存在自环或平行边。
给定两个数组 redEdges 和 blueEdges，其中：
redEdges[i] = [ai, bi] 表示图中存在一条从节点 ai 到节点 bi 的红色有向边，
blueEdges[j] = [uj, vj] 表示图中存在一条从节点 uj 到节点 vj 的蓝色有向边。
返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。
如果不存在这样的路径，那么 answer[x] = -1。
*/
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {

	type pair struct {
		to, color int
	}
	g := make([][]pair, n)

	for _, v := range redEdges {
		x, y := v[0], v[1]
		g[x] = append(g[x], pair{y, 0})
	}
	for _, v := range blueEdges {
		x, y := v[0], v[1]
		g[x] = append(g[x], pair{y, 1})
	}

	dis := make([][]int, n)

	for i := range dis {
		dis[i] = make([]int, 2)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = 0
	dis[0][1] = 0

	var dfs func(int, int, int)
	dfs = func(x, color, deep int) {
		if dis[x][color] < deep {
			return
		}
		dis[x][color] = deep
		for _, v := range g[x] {
			if color == v.color {
				continue
			}

			dfs(v.to, v.color, deep+1)
		}
	}

	dfs(0, 0, 0)
	dfs(0, 1, 0)
	ans := make([]int, n)
	for i := range dis {
		ans[i] = min(dis[i][0], dis[i][1])
		if ans[i] == math.MaxInt {
			ans[i] = -1
		}
	}

	return ans
}

func main() {
	fmt.Println(shortestAlternatingPaths(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
}
