package main

/*在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。*/
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	ans := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}

	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			ans = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {

				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				ans = max(ans, dp[i][j])
			} else {
				dp[i][j] = 0
			}
		}
	}
	return ans * ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
