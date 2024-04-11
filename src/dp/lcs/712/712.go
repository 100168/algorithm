package main

func minimumDeleteSum(word1 string, word2 string) int {

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
			s := 0
			for _, v := range word2[:j+1] {
				s += int(v)
			}
			return s
		}
		if j < 0 {
			s := 0
			for _, v := range word1[:i+1] {
				s += int(v)
			}
			return s
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}

		cur := 0
		if word1[i] == word2[j] {
			cur = dfs(i-1, j-1)
		} else {
			cur = min(dfs(i-1, j)+int(word1[i]), dfs(i, j-1)+int(word2[j]))
		}
		cache[i][j] = cur
		return cur
	}
	return dfs(m-1, n-1)
}
