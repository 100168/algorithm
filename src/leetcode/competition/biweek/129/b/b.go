package main

func numberOfRightTriangles(grid [][]int) int64 {

	m := len(grid)
	n := len(grid[0])

	sumR := make([]int, m)
	sumC := make([]int, n)

	for i := 0; i < m; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += grid[i][j]
		}
		sumR[i] = s
	}

	for j := 0; j < n; j++ {
		s := 0
		for i := 0; i < m; i++ {
			s += grid[i][j]
		}
		sumC[j] = s
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				col := sumR[i] - 1
				row := sumC[j] - 1
				ans += col * row
			}
		}
	}

	return int64(ans)
}

func main() {
	println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}))
}
