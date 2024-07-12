package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumMinimumPath(grid [][]int) int {

	l, r := 0, grid[0][0]

	for l <= r {
		m := (r + l) / 2
		if check(grid, m) {
			l = m + 1
		} else {
			r = m - 1
		}

	}
	return r

}

func check(gird [][]int, t int) bool {

	m := len(gird)
	n := len(gird[0])
	st := make([]pair, 0)
	st = append(st, pair{0, 0})

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	visited[0][0] = true
	for len(st) > 0 {
		temp := st
		st = nil

		for _, v := range temp {
			for _, d := range dirs {
				nx, ny := v.x+d.x, v.y+d.y
				if nx >= 0 && nx < m && ny >= 0 && ny < n && gird[nx][ny] >= t && !visited[nx][ny] {
					visited[nx][ny] = true
					if nx == m-1 && ny == n-1 {
						return true
					}
					st = append(st, pair{nx, ny})
				}
			}
		}
	}
	return false
}
