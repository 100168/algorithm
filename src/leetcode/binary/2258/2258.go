package main

var dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type pair struct{ x, y int }

func maximumMinutes(grid [][]int) int {

	m := len(grid)
	n := len(grid)
	l, r := 0, m*n-1

	for l <= r {
		m := (l + r) / 2
		if check(grid, m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if r < m*n {
		return r
	}

	return 1_000_000_000
}

func check(grid [][]int, t int) bool {

	m := len(grid)
	n := len(grid[0])
	f := make([]pair, 0)
	onFire := make([][]bool, len(grid))
	for i := range onFire {
		onFire[i] = make([]bool, len(grid[0]))
	}
	for i, v := range grid {
		for j := range v {
			if v[j] == 1 {
				f = append(f, pair{i, j})
				onFire[i][j] = true
			}
		}
	}

	var spreadFire func()
	spreadFire = func() {
		temp := f
		f = nil
		for j := range temp {
			cur := temp[j]
			for _, v := range dir {
				nx, ny := cur.x+v[0], cur.y+v[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if !onFire[nx][ny] && grid[nx][ny] == 0 {
						f = append(f, pair{nx, ny})
						onFire[nx][ny] = true
					}
				}
			}
		}
	}
	for ; t > 0 && len(f) > 0; t-- {
		spreadFire()
	}

	if onFire[0][0] {
		return false
	}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	walk := make([]pair, 0)
	walk = append(walk, pair{0, 0})
	vis[0][0] = true
	for len(walk) > 0 {
		temp := walk
		walk = nil
		for j := range temp {
			cur := temp[j]
			if onFire[cur.x][cur.y] {
				continue
			}
			for _, v := range dir {
				nx, ny := cur.x+v[0], cur.y+v[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] && !onFire[nx][ny] && grid[nx][ny] == 0 {
					vis[nx][ny] = true

					if nx == m-1 && ny == n-1 {
						return true
					}
					if grid[nx][ny] == 0 {
						walk = append(walk, pair{nx, ny})
					}
				}
			}
		}
		spreadFire()
	}
	return false

}
