package main

import (
	"fmt"
)

/*
*
给你一个大小为 m x n ，由若干 0 和 1 组成的二维网格 grid ，其中 1 表示陆地， 0 表示水。
岛屿 由水平方向或竖直方向上相邻的 1 （陆地）连接形成。

如果 恰好只有一座岛屿 ，则认为陆地是 连通的 ；否则，陆地就是 分离的 。

一天内，可以将 任何单个 陆地单元（1）更改为水单元（0）。

返回使陆地分离的最少天数。
*/
func minDays2(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(int, int)

	dfs = func(x, y int) {
		grid[x][y] = -1
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] == 1 {
				dfs(nx, ny)
			}
		}
	}

	count := func() int {
		cnt := 0
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					dfs(i, j)
					cnt++
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == -1 {
					grid[i][j] = 1
				}
			}
		}
		return cnt
	}

	if count() != 1 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if count() != 1 {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}
	return 2
}

func minDays(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])
	cntOne := 0

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	g := make(map[int][]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cntOne++
				for _, d := range dirs {
					nx, ny := i+d[0], j+d[1]
					if nx < 0 || nx >= m || ny < 0 || ny >= n {
						continue
					}
					if grid[nx][ny] == 1 {
						g[i*n+j] = append(g[i*n+j], nx*n+ny)
					}
				}
			}
		}
	}
	//特殊处理
	if cntOne == 1 {
		return 1
	}
	if len(g) == 0 {
		return 0
	}

	isCut := make([]bool, m*n)
	dfn := make([]int, m*n) // DFS 到结点 v 的时间（从 1 开始）
	low := make([]int, m*n) //v节点最小时间 的时间（从 1 开始）
	// low[v] 定义为以下两种情况的最小值
	// 1. dfn[v]
	// 2. subtree(v) 的返祖边所指向的节点的 dfn，也就是经过恰好一条不在 DFS 树上的边，能够到达 subtree(v) 的节点的 dfn
	dfsClock := 0
	cnt := 0
	var tarjan func(int, int) int
	tarjan = func(v, fa int) int { // 无需考虑重边
		dfsClock++
		dfn[v] = dfsClock
		low[v] = dfsClock
		childCnt := 0
		for _, w := range g[v] {
			if dfn[w] == 0 {
				childCnt++
				lowW := tarjan(w, v)
				low[v] = min(low[v], lowW)
				if lowW >= dfn[v] {
					isCut[v] = true
					cnt++
				}
			} else {
				low[v] = min(low[v], dfn[w])
			}
		}
		if fa == -1 && childCnt == 1 {
			if isCut[v] {
				cnt--
			}
			isCut[v] = false
		}
		return low[v]
	}

	cntR := 0
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 && dfn[i*n+j] == 0 {
				if cntR > 0 {
					return 0
				}
				cntR++
				tarjan(i*n+j, -1)
			}
		}
	}
	if cnt > 0 {
		return 1
	}
	return 2
}

//{{1, 1, 0, 1, 1},
//{1, 1, 1, 1, 1},
//{1, 1, 0, 1, 1},
//{1, 1, 0, 1, 1}}

//[[1,1,0,1,1],
//[1,1,1,1,1],
//[1,1,0,1,1],
//[1,1,1,1,1]]

//[[1,0,1,1,1],
//[0,1,1,1,1],
//[0,1,1,1,1],
//[1,1,1,1,1],
//[1,0,1,0,1],
//[0,1,1,1,1],
//[0,1,0,0,1],
//[1,1,1,0,1]]

func main() {
	//fmt.Println(minDays([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
	fmt.Println(minDays([][]int{{1, 0}, {1, 1}}))
	//fmt.Println(minDays([][]int{{1, 0, 1, 1, 1}, {0, 1, 1, 1, 1}, {0, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}, {1, 1, 1, 0, 1}}))
}
