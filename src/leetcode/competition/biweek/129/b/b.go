package main

/*
*给你一个二维 boolean 矩阵 grid 。

请你返回使用 grid 中的 3 个元素可以构建的 直角三角形 数目，且满足 3 个元素值 都 为 1 。
注意：
如果 grid 中 3 个元素满足：一个元素与另一个元素在 同一行，同时与第三个元素在 同一列 ，那么这 3 个元素称为一个 直角三角形 。这 3 个元素互相之间不需要相邻。
*/
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
