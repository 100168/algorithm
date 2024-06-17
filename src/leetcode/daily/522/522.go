package main

import (
	"fmt"
)

func findLUSlength(strs []string) int {

	ans := -1
	for i, v1 := range strs {

		flag := false
		for j, v2 := range strs {
			if flag {
				break
			}
			if j == i {
				continue
			}
			if check(v1, v2) {
				flag = true
			}

		}
		if !flag {
			ans = max(ans, len(v1))
		}
	}
	return ans
}

func check(s1, s2 string) bool {

	m := len(s1)
	n := len(s2)
	if len(s1) > len(s2) {
		return false
	}

	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n] == m
}

func main() {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
}
