package main

func minimumArea(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	up := m - 1
	down := 0
	left := n - 1
	right := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				up = min(up, i)
				down = max(down, i)
				left = min(left, j)
				right = max(right, j)
			}
		}
	}

	return (down - up + 1) * (right - left + 1)

}
