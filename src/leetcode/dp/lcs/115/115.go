package main

func numDistinct(s string, t string) int {

	mod := 1_000_000_007

	m := len(s)
	n := len(t)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, j int) int {
		if j < 0 {
			return 1
		}
		if i < 0 {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		cur := 0
		if s[i] == t[j] {
			cur += dfs(i-1, j-1)
			cur %= mod
		}
		cur += dfs(i-1, j)
		cur %= mod
		cache[i][j] = cur
		return cur
	}

	return dfs(m-1, n-1)
}

func main() {
	numDistinct("aaabc", "abc")
}
