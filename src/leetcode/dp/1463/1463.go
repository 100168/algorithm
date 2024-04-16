package main

func cherryPickup(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)

	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)

			for k := 0; k < n; k++ {
				dp[i][j][k] = -7001
			}
		}

	}

	dp[0][0][n-1] = grid[0][0] + grid[0][n-1]

	ans := -1
	dir := []int{-1, 0, 1}
	for i := 1; i < m; i++ {
		for j1 := 0; j1 < n; j1++ {
			for j2 := j1 + 1; j2 < n; j2++ {
				res := -7001
				for _, d1 := range dir {
					if j1+d1 >= 0 && j1+d1 < n {
						for _, d2 := range dir {
							if j2+d2 >= 0 && j2+d2 < n {
								res = max(dp[i-1][j1+d1][j2+d2], res)
							}
						}
					}
				}
				res += grid[i][j1] + grid[i][j2]
				dp[i][j1][j2] = res
				if i == m-1 {
					ans = max(ans, dp[i][j1][j2])
				}
			}
		}

	}
	return ans
}
