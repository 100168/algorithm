package main

import (
	"math"
)

// 多解，并查集，最短路，二分
func swimInWater(grid [][]int) int {

	l, r := 0, math.MaxInt

	for l <= r {
		m := (r-l)/2 + l

		if check(m, grid) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func check(t int, grid [][]int) bool {

	type pair struct{ x, y int }

	m := len(grid)
	n := len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	visit[0][0] = true
	s := make([]pair, 0)
	s = append(s, pair{0, 0})
	for len(s) > 0 {
		pop := s[0]
		s = s[1:]
		if pop.x == m-1 && pop.y == n-1 {
			return true
		}
		for _, dir := range dirs {
			x, y := pop.x+dir[0], pop.y+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && !visit[x][y] && t >= grid[x][y] {
				s = append(s, pair{x, y})
				visit[x][y] = true
			}
		}

	}
	return false
}

func main() {

	println(swimInWater([][]int{{0, 2}, {1, 3}}))
}
