package main

import (
	"fmt"
	"math"
)

/*
*给你一个正整数 n ，表示总共有 n 个城市，城市从 1 到 n 编号。
给你一个二维数组 roads ，其中 roads[i] = [ai, bi, distancei] 表示城市 ai 和 bi 之间有一条 双向 道路，
道路距离为 distancei 。城市构成的图不一定是连通的。

两个城市之间一条路径的 分数 定义为这条路径中道路的 最小 距离。
城市 1 和城市 n 之间的所有路径的 最小 分数。
注意：
一条路径指的是两个城市之间的道路序列。
一条路径可以 多次 包含同一条道路，你也可以沿着路径多次到达城市 1 和城市 n 。
测试数据保证城市 1 和城市n 之间 至少 有一条路径。
*/
func minScore(n int, roads [][]int) int {

	g := make([][]int, n)

	ds := make([]int, n)

	for i := range ds {
		ds[i] = math.MaxInt
	}
	for _, v := range roads {
		x, y, d := v[0]-1, v[1]-1, v[2]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		ds[x] = min(ds[x], d)
		ds[y] = min(ds[y], d)
	}

	visited := make([]bool, n)
	var dfs func(int)

	ans := math.MaxInt
	dfs = func(x int) {
		visited[x] = true
		ans = min(ans, ds[x])
		for _, y := range g[x] {
			if !visited[y] {
				dfs(y)
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(minScore(4, [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}))
}
