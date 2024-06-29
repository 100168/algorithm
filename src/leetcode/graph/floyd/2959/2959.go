package main

import (
	"fmt"
	"math"
)

/*
*一个公司在全国有 n 个分部，它们之间有的有道路连接。一开始，所有分部通过这些道路两两之间互相可以到达。

公司意识到在分部之间旅行花费了太多时间，所以它们决定关闭一些分部（也可能不关闭任何分部），
同时保证剩下的分部之间两两互相可以到达且最远距离不超过 maxDistance 。

两个分部之间的 距离 是通过道路长度之和的 最小值 。

给你整数 n ，maxDistance 和下标从 0 开始的二维整数数组 roads ，其中 roads[i] = [ui, vi, wi] 表示一条从 ui 到 vi 长度为 wi的 无向 道路。

请你返回关闭分部的可行方案数目，满足每个方案里剩余分部之间的最远距离不超过 maxDistance。

注意，关闭一个分部后，与之相连的所有道路不可通行。

注意，两个分部之间可能会有多条道路。
*/
func numberOfSets(n int, maxDistance int, roads [][]int) int {

	ans := 0
	w := make([][]int, n)
	for j := range w {
		w[j] = make([]int, n)
		for k := range w[j] {
			w[j][k] = math.MaxInt / 2
		}
		w[j][j] = 0
	}
	for _, v := range roads {
		x, y, d := v[0], v[1], v[2]
		w[x][y] = min(w[x][y], d)
		w[y][x] = min(w[x][y], d)
	}
	f := make([][]int, n)

	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for i := 0; i < 1<<n; i++ {

		for j, r := range w {
			if 1<<j&i == 0 {
				continue
			}
			copy(f[j], r)
		}
		for k := 0; k < n; k++ {
			if 1<<k&i == 0 {
				continue
			}
			for x := 0; x < n; x++ {
				if 1<<x&i == 0 {
					continue
				}
				for y := 0; y < n; y++ {
					f[x][y] = min(f[x][y], f[x][k]+f[k][y])
				}
			}
		}
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				if 1<<x&i == 0 || 1<<y&i == 0 {
					continue
				}
				if f[x][y] > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(numberOfSets(3, 12, [][]int{{1, 0, 11}, {1, 0, 16}, {0, 2, 13}}))
}
