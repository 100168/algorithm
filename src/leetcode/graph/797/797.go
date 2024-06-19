package main

import "slices"

/*
*
给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）

	graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。
*/
func allPathsSourceTarget(graph [][]int) [][]int {

	n := len(graph)
	var ans [][]int

	var dfs func(int, int, []int)

	dfs = func(x int, fa int, path []int) {
		if x == n-1 {
			ans = append(ans, slices.Clone(path))
			return
		}

		for _, y := range graph[x] {
			if y == fa {
				continue
			}
			path = append(path, y)
			dfs(y, x, path)
			path = path[:len(path)-1]
		}

	}

	dfs(0, -1, []int{0})

	return ans
}
