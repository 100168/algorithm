package main

func canMakeSquare(grid [][]byte) bool {

	n := len(grid)

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {

			cntB := 0
			cntW := 0

			if grid[i][j] == 'B' {
				cntB++
			} else {
				cntW++
			}
			if grid[i][j-1] == 'B' {
				cntB++
			} else {
				cntW++
			}
			if grid[i-1][j-1] == 'B' {
				cntB++
			} else {
				cntW++
			}
			if grid[i-1][j] == 'B' {
				cntB++
			} else {
				cntW++
			}

			if cntB >= 3 || cntW >= 3 {
				return true
			}

		}
	}
	return false
}
