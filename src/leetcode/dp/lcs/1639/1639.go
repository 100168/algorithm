package main

func numWays(words []string, target string) int {

	mod := 1_000_000_007
	n := len(words[0])

	cnt := make([][]int, n)

	for i := range cnt {
		cnt[i] = make([]int, 26)
	}

	for _, word := range words {
		for i := range word {
			cnt[i][word[i]-'a']++
		}
	}

	cache := make([][]int, n)

	t := len(target)
	for i := range cache {
		cache[i] = make([]int, t)
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
		curT := target[j] - 'a'
		if cnt[i][curT] > 0 {
			cur += dfs(i-1, j-1) * cnt[i][curT]
			cur %= mod
		}

		cur += dfs(i-1, j)
		cur %= mod
		cache[i][j] = cur
		return cur
	}
	return dfs(n-1, t-1)
}
