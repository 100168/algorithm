package main

func isInterleave(s1 string, s2 string, s3 string) bool {

	m := len(s1)
	n := len(s2)
	k := len(s3)
	if m+n != k {
		return false
	}

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) bool

	dfs = func(i, j int) bool {
		if i < 0 {
			for t, v := range s2[:j+1] {
				if uint8(v) != s3[t] {
					return false
				}
			}
			return true

		}
		if j < 0 {
			for t, v := range s1[:i+1] {
				if uint8(v) != s3[t] {
					return false
				}
			}
			return true
		}

		if cache[i][j] != -1 {
			if cache[i][j] == 1 {
				return true
			} else {
				return false
			}
		}
		index := m - (i + 1) + n - (j + 1) - 1
		cur := false
		if s1[i] == s3[index] {
			cur = dfs(i-1, j)
		}
		if s2[j] == s3[index] {
			cur = cur || dfs(i, j-1)
		}

		if cur {
			cache[i][j] = 1
		} else {
			cache[i][j] = 0
		}
		return cur

	}
	return dfs(m-1, n-1)

}
