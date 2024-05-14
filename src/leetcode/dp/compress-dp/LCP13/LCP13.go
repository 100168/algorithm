package main

import (
	"fmt"
	"math"
)

var (
	dx   = []int{1, -1, 0, 0}
	dy   = []int{0, 0, 1, -1}
	n, m int
)

func minimalSteps(maze []string) int {
	n, m = len(maze), len(maze[0])
	// 机关 & 石头
	var buttons, stones [][]int
	// 起点 & 终点
	sx, sy, tx, ty := -1, -1, -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch maze[i][j] {
			case 'M':
				buttons = append(buttons, []int{i, j})
			case 'O':
				stones = append(stones, []int{i, j})
			case 'S':
				sx, sy = i, j
			case 'T':
				tx, ty = i, j
			}
		}
	}

	nb, ns := len(buttons), len(stones)
	startDist := bfs(sx, sy, maze)
	// 边界情况：没有机关
	if nb == 0 {
		return startDist[tx][ty]
	}
	// 从某个机关到其他机关 / 起点与终点的最短距离。
	dist := make([][]int, nb)
	for i := 0; i < nb; i++ {
		dist[i] = make([]int, nb+2)
		for j := 0; j < nb+2; j++ {
			dist[i][j] = -1
		}
	}
	// 中间结果
	dd := make([][][]int, nb)
	for i := 0; i < nb; i++ {
		dd[i] = bfs(buttons[i][0], buttons[i][1], maze)
		// 从某个点到终点不需要拿石头
		dist[i][nb+1] = dd[i][tx][ty]
	}
	for i := 0; i < nb; i++ {
		tmp := -1
		for k := 0; k < ns; k++ {
			midX, midY := stones[k][0], stones[k][1]
			if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
				if tmp == -1 || tmp > dd[i][midX][midY]+startDist[midX][midY] {
					tmp = dd[i][midX][midY] + startDist[midX][midY]
				}
			}
		}
		dist[i][nb] = tmp
		for j := i + 1; j < nb; j++ {
			mn := -1
			for k := 0; k < ns; k++ {
				midX, midY := stones[k][0], stones[k][1]
				if dd[i][midX][midY] != -1 && startDist[midX][midY] != -1 {
					if mn == -1 || mn > dd[i][midX][midY]+dd[j][midX][midY] {
						mn = dd[i][midX][midY] + dd[j][midX][midY]
					}
				}
			}
			dist[i][j] = mn
			dist[j][i] = mn
		}
	}
	// 无法达成的情形
	for i := 0; i < nb; i++ {
		if dist[i][nb] == -1 || dist[i][nb+1] == -1 {
			return -1
		}
	}
	// dp 数组， -1 代表没有遍历到
	dp := make([][]int, 1<<nb)
	for i := 0; i < (1 << nb); i++ {
		dp[i] = make([]int, nb)
		for j := 0; j < nb; j++ {
			dp[i][j] = -1
		}
	}
	for i := 0; i < nb; i++ {
		dp[1<<i][i] = dist[i][nb]
	}

	memo := make([][]int, 1<<nb)
	for i := range memo {
		memo[i] = make([]int, nb)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(mask int, pre int) int {
		if mask == 1<<nb-1 {
			return dist[pre][nb+1]
		}

		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}

		cur := math.MaxInt
		for j := 0; j < nb; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask, j)+dist[pre][j])
			}
		}
		memo[mask][pre] = cur
		if cur == 0 {
			println(mask, pre)
		}
		return cur
	}
	res := math.MaxInt

	for i := 0; i < nb; i++ {
		res = min(res, dist[i][nb]+dfs(1<<i, i))
	}
	return res
}

func bfs(x, y int, maze []string) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
		for j := 0; j < m; j++ {
			ret[i][j] = -1
		}
	}
	ret[x][y] = 0
	var queue [][]int
	queue = append(queue, []int{x, y})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		curX, curY := p[0], p[1]
		for k := 0; k < 4; k++ {
			nx, ny := curX+dx[k], curY+dy[k]
			if inBound(nx, ny) && maze[nx][ny] != '#' && ret[nx][ny] == -1 {
				ret[nx][ny] = ret[curX][curY] + 1
				queue = append(queue, []int{nx, ny})
			}
		}
	}
	return ret
}

func inBound(x, y int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}

func main() {
	fmt.Println(minimalSteps([]string{
		"MMMMM",
		"MS#MM",
		"MM#TO"}))
}
