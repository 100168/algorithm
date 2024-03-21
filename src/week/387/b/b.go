package main

func countSubmatrices(grid [][]int, k int) int {

	m := len(grid)
	n := len(grid[0])

	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = grid[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
			if sum[i][j] == k {
				ans++
			}

		}
	}
	return ans
}
