package main

import "slices"

func maxPoints(points [][]int) int64 {

	m := len(points)
	n := len(points[0])
	dp := make([][]int64, 2)

	dpLeft := make([]int64, n)
	dpRight := make([]int64, n)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int64, n)
		for j := 0; j < n; j++ {
			dp[0][j] = int64(points[0][j])
		}
	}

	dpLeft[0] = int64(points[0][0])
	dpRight[n-1] = int64(points[0][n-1])
	for j := 1; j < n; j++ {
		dpLeft[j] = max(int64(points[0][j]), dpLeft[j-1]-1)
	}
	for j := n - 2; j >= 0; j-- {
		dpRight[j] = max(int64(points[0][j]), dpRight[j+1]-1)
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			preMax := max(dpLeft[j], dpRight[j])
			dp[i%2][j] = preMax + int64(points[i][j])
		}
		dpLeft[0] = dp[i%2][0]
		for j := 1; j < n; j++ {
			dpLeft[j] = max(dpLeft[j-1]-1, dp[i%2][j])
		}
		dpRight[n-1] = dp[i%2][n-1]
		for j := n - 2; j >= 0; j-- {
			dpRight[j] = max(dp[i%2][j], dpRight[j+1]-1)
		}
	}
	return slices.Max(dp[(m-1)%2])
}

func main() {
	println(maxPoints([][]int{{4, 1}, {5, 4}, {4, 3}, {1, 3}}))
}
