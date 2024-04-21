package main

import "strconv"

func numberOfArrays(s string, k int) int {

	n := len(s)

	ks := strconv.Itoa(k)
	lk := len(ks)
	memo := make([]int, n)

	mod := 1_000_000_007

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {

		if i < 0 {
			return 1
		}

		if memo[i] != -1 {
			return memo[i]
		}

		cur := 0
		for j := i - 1; (i-j) < lk && j >= -1; j-- {
			if s[j+1] == '0' {
				continue
			}
			cur += dfs(j)
		}

		if i+1 >= lk && s[i+1-lk:i+1] <= ks && s[i+1-lk] != '0' {
			cur += dfs(i - lk)
		}
		memo[i] = cur % mod

		return memo[i]

	}
	return dfs(n - 1)
}
