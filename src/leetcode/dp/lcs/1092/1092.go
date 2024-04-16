package main

func shortestCommonSupersequence(str1 string, str2 string) string {

	m := len(str1)
	n := len(str2)

	dp := make([][]string, m+1)
	for i := range dp {
		dp[i] = make([]string, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(str1[i-1])
			} else {
				s1 := dp[i-1][j] + string(str1[i-1])
				s2 := dp[i][j-1] + string(str2[j-1])
				if len(s1) < len(s2) {
					dp[i][j] = s1
				} else {
					dp[i][j] = s2
				}
			}
		}
	}
	return dp[m][n]
}

func shortestCommonSupersequence2(str1 string, str2 string) string {

	m := len(str1)
	n := len(str2)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i < 0 {
			return j
		}
		if j < 0 {
			return i
		}
		if cache[i][j] != -1 {
			return cache[i][j]
		}
		if str1[i] == str2[j] {
			cache[i][j] = dfs(i-1, j-1) + 1
			return cache[i][j]
		}
		cur := min(dfs(i-1, j), dfs(i, j-1)) + 1
		cache[i][j] = cur
		return cache[i][j]
	}
	var makeAns func(int, int) string
	makeAns = func(i, j int) string {
		if i < 0 { // s 是空串，返回剩余的 t
			return str2[:j+1]
		}
		if j < 0 { // t 是空串，返回剩余的 s
			return str1[:i+1]
		}
		if str1[i] == str2[j] { // 最短公共超序列一定包含 s[i]
			return makeAns(i-1, j-1) + string(str2[i])
		}
		// 如果下面 if 成立，说明上面 dfs 中的 min 取的是 dfs(i - 1, j)
		// 说明 dfs(i - 1, j) 对应的公共超序列更短
		// 那么就在 makeAns(i - 1, j) 的结果后面加上 s[i]
		// 否则说明 dfs(i, j - 1) 对应的公共超序列更短
		// 那么就在 makeAns(i, j - 1) 的结果后面加上 t[j]
		if dfs(i, j) == dfs(i-1, j)+1 {
			return makeAns(i-1, j) + string(str1[i])
		}
		return makeAns(i, j-1) + string(str2[j])
	}
	return makeAns(n-1, m-1)
}
