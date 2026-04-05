package main

import (
	"math"
)

func minOperations(s1 string, s2 string, x int) int {

	n := len(s1)

	diff := make([]bool, n)

	cur := uint8(0)
	for i := 0; i < n; i++ {
		cur ^= s1[i] ^ s2[i]
		diff[i] = s1[i] != s2[i]
	}

	if cur > 0 {
		return -1
	}

	memo := make([][]map[bool]int, n)

	for i := range memo {
		memo[i] = make([]map[bool]int, n)

		for j := range memo[i] {
			memo[i][j] = make(map[bool]int)
		}
	}

	var dfs func(int, int, bool) int

	dfs = func(i, j int, pre bool) int {
		if i < 0 {
			if pre || j > 0 {
				return math.MaxInt / 2
			}
			return 0
		}

		if _, ok := memo[i][j][pre]; ok {
			return memo[i][j][pre]
		}
		cur := math.MaxInt / 2

		if pre == diff[i] {
			cur = dfs(i-1, j, false)
		} else {
			cur = dfs(i-1, j+1, false) + x
			cur = min(cur, dfs(i-1, j, true)+1)
			if j > 0 {
				cur = min(cur, dfs(i-1, j-1, false))
			}
		}
		memo[i][j][pre] = cur
		return cur
	}

	return dfs(n-1, 0, false)

}

func main() {
	println(minOperations("101101", "000000", 6))
}
