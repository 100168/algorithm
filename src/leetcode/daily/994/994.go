package main

func orangesRotting(grid [][]int) int {

	type pair struct{ x, y int }
	var st []pair
	m, n := len(grid), len(grid[0])

	cntFresh := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				st = append(st, pair{i, j})
			} else if grid[i][j] == 1 {
				cntFresh++
			}
		}
	}

	dir := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	ans := -1
	for len(st) > 0 {
		ans++
		temp := st
		st = nil
		for _, v := range temp {
			for _, d := range dir {
				x := v.x + d[0]
				y := v.y + d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					st = append(st, pair{x, y})
					grid[x][y] = 2
					cntFresh--
				}
			}
		}

	}

	if cntFresh > 0 {
		return -1
	}

	return max(ans, 0)
}
