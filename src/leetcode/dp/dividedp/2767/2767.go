package main

import (
	"math"
	"strconv"
)

func minimumBeautifulSubstrings(s string) int {

	n := len(s)
	dict := make(map[string]bool)
	base := uint(5)
	cur := uint(1)
	for i := 0; ; i++ {
		str := strconv.FormatInt(int64(cur), 2)
		if len(str) > n {
			break
		}
		dict[str] = true
		cur *= base
	}

	memo := make([]int, n+1)

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(i int) int {
		if i == 0 {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		cur := math.MaxInt / 2

		for j := i - 1; j >= 0; j-- {
			str := s[j:i]
			if dict[str] {
				cur = min(cur, dfs(j)+1)
			}
		}

		memo[i] = cur
		return cur
	}

	ans := dfs(n)
	if ans > n {
		return -1
	}
	return ans
}
func main() {

	println(minimumBeautifulSubstrings("1011"))
}
