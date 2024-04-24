package main

import "math"

/*
*
给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目
。网格图中总共恰好有 9 个石头，一个格子里可能会有 多个 石头。

每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。

请你返回每个格子恰好有一个石头的 最少移动次数 。
*/
type pair struct{ x, y int }

func minimumMoves(grid [][]int) int {
	var from, to []pair
	for i, row := range grid {
		for j, cnt := range row {
			if cnt > 1 {
				for k := 1; k < cnt; k++ {
					from = append(from, pair{i, j})
				}
			} else if cnt == 0 {
				to = append(to, pair{i, j})
			}
		}
	}

	ans := math.MaxInt

	var dfs func(int, int, []pair)

	dfs = func(i, mask int, path []pair) {
		if i < 0 {
			s := 0
			for i, v := range path {
				s += abs(from[i].x-v.x) + abs(from[i].y-v.y)
			}
			ans = min(ans, s)
		}
		for j := 0; j < len(to); j++ {
			if 1<<j&mask == 0 {
				path := append(path, to[j])
				dfs(i-1, mask|1<<j, path)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(len(to)-1, 0, []pair{})
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
