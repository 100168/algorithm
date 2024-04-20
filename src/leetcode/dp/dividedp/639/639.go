package main

func numDecodings(s string) int {

	n := len(s)

	mod := 1_000_000_007
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

		if s[i] == '*' || s[i] > '0' {
			if s[i] == '*' {
				cur += 9 * dfs(i-1)
			} else {
				cur += dfs(i - 1)
			}

		}

		if i > 0 {
			if s[i-1] == '1' && s[i] != '*' || (s[i-1] == '2' && s[i] <= '6' && s[i] != '*') {
				cur += dfs(i - 2)
			}

			if s[i-1] == '*' {
				if s[i] == '*' {
					cur += 15 * dfs(i-2)
				} else {
					cur += dfs(i - 2)
					if s[i] <= '6' {
						cur += dfs(i - 2)
					}
				}
			} else if s[i] == '*' {
				if s[i-1] == '1' {
					cur += 9 * dfs(i-2)
				} else if s[i-1] == '2' {
					cur += 6 * dfs(i-2)
				}
			}
		}

		memo[i] = cur % mod
		return memo[i]
	}

	return dfs(n - 1)
}
