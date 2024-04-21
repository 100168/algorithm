package main

import "strconv"

func crackNumber(crackNumber int) int {
	s := strconv.Itoa(crackNumber)
	n := len(s)

	dp := make([]int, n+1)

	dp[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i > 1 && (s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '5') {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]

}

func main() {
	println(crackNumber(216612))
}
