package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maximumSafenessFactor(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = -1
		}
	}

	st := make([]pair, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				st = append(st, pair{i, j})
				dis[i][j] = 0
			}
		}
	}
	for len(st) > 0 {
		temp := st
		st = nil
		for _, v := range temp {
			for _, d := range dirs {
				x := v.x + d.x
				y := v.y + d.y
				if x >= 0 && x < m && y >= 0 && y < n && dis[x][y] == -1 {
					dis[x][y] = dis[v.x][v.y] + 1
					st = append(st, pair{x, y})
				}
			}
		}
	}
	l, r := 1, dis[0][0]
	for l <= r {
		m := (r-l)/2 + l
		if check(m, dis) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
func check(t int, dis [][]int) bool {
	path := make([]pair, 0)
	path = append(path, pair{0, 0})
	m := len(dis)
	n := len(dis[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	visit[0][0] = true
	for len(path) > 0 {
		temp := path
		path = nil
		for _, p := range temp {
			if p.x == m-1 && p.y == n-1 {
				return true
			}
			for _, dir := range dirs {
				x := p.x + dir.x
				y := p.y + dir.y

				if x >= 0 && y >= 0 && x < m && y < n && !visit[x][y] && dis[x][y] >= t {
					visit[x][y] = true
					path = append(path, pair{x, y})
				}
			}

		}
	}
	return false

}
func main() {
	println(maximumSafenessFactor([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
}
