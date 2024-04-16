package main

func minDistance(word1 string, word2 string) int {

	m := len(word1)
	n := len(word2)
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {

		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		cur := 0
		if word1[i] == word2[j] {
			cur = dfs(i-1, j-1)
		} else {
			cur = min(dfs(i-1, j), dfs(i, j-1)) + 1
		}
		cache[i][j] = cur
		return cur
	}

	return dfs(m-1, n-1)
}
