package main

func numDecodings(s string) int {

	n := len(s)

	memo := make([]int, n)

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

		if s[i] > '0' {
			cur += dfs(i - 1)
		}
		if i-1 >= 0 {
			if s[i-1] == '1' || s[i-1] > '0' && s[i-1] <= '2' && s[i] <= '6' {
				cur += dfs(i - 2)
			}

		}

		memo[i] = cur
		return cur

	}

	return dfs(n - 1)
}
