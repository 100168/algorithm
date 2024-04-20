package main

import "strings"

func minExtraChar(s string, dictionary []string) int {

	n := len(s)

	memo := make([]int, n+1)

	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {

		if i <= 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		cur := i

		for _, v := range dictionary {

			if i >= len(v) {
				index := strings.LastIndex(s[:i], v)
				if index == -1 {
					continue
				}
				diff := i - len(v) - index
				cur = min(cur, dfs(index)+diff)
			}
		}
		memo[i] = cur
		return cur
	}

	return dfs(n)
}
