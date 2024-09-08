package main

import (
	"fmt"
	"math"
	"slices"
)

/**
给你一个 50 x 50 的国际象棋棋盘，棋盘上有 一个 马和一些兵。给你两个整数 kx 和 ky ，其中 (kx, ky) 表示马所在的位置，同时还有一个二维数组 positions ，其中 positions[i] = [xi, yi] 表示第 i 个兵在棋盘上的位置。

Alice 和 Bob 玩一个回合制游戏，Alice 先手。玩家的一次操作中，可以执行以下操作：

玩家选择一个仍然在棋盘上的兵，然后移动马，通过 最少 的 步数 吃掉这个兵。注意 ，玩家可以选择 任意 一个兵，不一定 要选择从马的位置出发 最少 移动步数的兵。
在马吃兵的过程中，马 可能 会经过一些其他兵的位置，但这些兵 不会 被吃掉。只有 选中的兵在这个回合中被吃掉。
Alice 的目标是 最大化 两名玩家的 总 移动次数，直到棋盘上不再存在兵，而 Bob 的目标是 最小化 总移动次数。

假设两名玩家都采用 最优 策略，请你返回 Alice 可以达到的 最大 总移动次数。

在一次 移动 中，如下图所示，马有 8 个可以移动到的位置，每个移动位置都是沿着坐标轴的一个方向前进 2 格，然后沿着垂直的方向前进 1 格。


*/

func maxMoves(kx int, ky int, positions [][]int) int {

	n := len(positions)
	dir := [][]int{{2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}}

	type pair struct{ x, y int }

	var findMin func(int, int, []int) int

	newPos := slices.Clone(positions)

	newPos = append(newPos, []int{kx, ky})
	findMin = func(x int, y int, pos []int) int {

		stack := []pair{{x, y}}

		visited := make([]bool, 50*50)
		cnt := 0
		for {

			temp := stack
			stack = nil
			for _, p := range temp {

				if p.x == pos[0] && p.y == pos[1] {
					return cnt
				}
				for _, d := range dir {
					nx, ny := p.x+d[0], p.y+d[1]
					if nx >= 0 && nx < 50 && ny >= 0 && ny < 50 && !visited[nx*50+ny] {
						visited[nx*50+ny] = true
						stack = append(stack, pair{nx, ny})
					}
				}
			}

			cnt++
		}
	}

	minPaths := make([][]int, n+1)

	for i := range minPaths {

		minPaths[i] = make([]int, n+1)
	}

	for i, p := range newPos {
		for j := i + 1; j <= n; j++ {
			minPaths[i][j] = findMin(p[0], p[1], newPos[j])
			minPaths[j][i] = minPaths[i][j]
		}
	}

	f := make([][]map[bool]int, 1<<n)

	for i := range f {
		f[i] = make([]map[bool]int, n+1)
		for j := range f[i] {
			f[i][j] = make(map[bool]int)
		}
	}

	var dfs func(int, int, bool) int

	dfs = func(mask int, pre int, first bool) int {

		if mask == 1<<n-1 {
			return 0
		}

		if _, ok := f[mask][pre][first]; ok {
			return f[mask][pre][first]
		}

		minPath := math.MaxInt / 2
		maxPath := 0
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				if first {
					maxPath = max(maxPath, dfs(1<<j|mask, j, false)+minPaths[pre][j])
				} else {
					minPath = min(minPath, dfs(1<<j|mask, j, true)+minPaths[pre][j])
				}
			}
		}
		if first {
			f[mask][pre][first] = maxPath
		} else {
			f[mask][pre][first] = minPath
		}
		return f[mask][pre][first]
	}
	return dfs(0, n, true)
}

func main() {
	//fmt.Println(maxMoves(1, 1, [][]int{{0, 0}}))
	fmt.Println(maxMoves(0, 2, [][]int{{1, 1}, {2, 2}, {3, 3}}))
}
