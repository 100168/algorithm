package main

import "fmt"

func isMatch(s string, p string) bool {

	m := len(s)
	n := len(p)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) bool

	dfs = func(i, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 {
			for v := range p[:j+1] {
				if p[v] != '*' {
					return false
				}
			}
			return true
		}
		if j < 0 {
			return false
		}
		if cache[i][j] != -1 {
			return cache[i][j] > 0
		}

		cur := false

		if p[j] == '*' {
			cur = dfs(i-1, j) || dfs(i, j-1)
		} else {
			if s[i] == p[j] || p[j] == '?' {
				cur = dfs(i-1, j-1)
			}
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

func main() {
	fmt.Println(isMatch("aa", "*"))
}
