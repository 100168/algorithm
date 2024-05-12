package main

import (
	"fmt"
	"math"
)

func minDays(n int) int {
	memo := make(map[int]int)
	var dfs func(int) int

	dfs = func(x int) int {
		if x == 0 {
			return 0
		}
		if x == 1 {
			return 1
		}

		if _, ok := memo[x]; ok {
			return memo[x]
		}
		cur := math.MaxInt

		cur = min(cur, dfs(x/3)+x%3+1)

		cur = min(cur, dfs(x/2)+x%2+1)

		memo[x] = cur
		return cur
	}
	return dfs(n)
}

func main() {
	fmt.Println(minDays(10))
}
