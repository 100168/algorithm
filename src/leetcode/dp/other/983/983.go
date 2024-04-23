package main

import "math"

func mincostTickets(days []int, costs []int) int {

	memo := make([]int, 366)
	for i := range memo {
		memo[i] = -1
	}

	dayM := map[int]bool{}
	for _, d := range days {
		dayM[d] = true
	}
	var dfs func(int) int

	dfs = func(i int) int {

		if i > 365 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		cur := math.MaxInt / 2
		if dayM[i] {
			cur = min(cur, dfs(i+1)+costs[0])
			cur = min(cur, dfs(i+7)+costs[1])
			cur = min(cur, dfs(i+30)+costs[2])
		} else {
			cur = dfs(i + 1)
		}

		memo[i] = cur
		return cur
	}

	return dfs(1)
}
