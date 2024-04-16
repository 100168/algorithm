package main

func jewelleryValue(frame [][]int) int {

	m := len(frame)
	n := len(frame[0])
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = frame[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 {
				dp[i&1][j] = dp[(i-1)&1][j] + frame[i][j]
			}
			if j-1 >= 0 {
				dp[i&1][j] = max(dp[i&1][j-1]+frame[i][j], dp[i&1][j])
			}
		}
	}
	return dp[(m-1)&1][n-1]
}

func main() {
	println(jewelleryValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}))
}
