package _012

import "strconv"

//给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。

func numDupDigitsAtMostN(n int) int {

	s := strconv.Itoa(n)

	l := len(s)
	m := 1024

	memo := make([][]int, l)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return n - dfs(0, 0, true, false, s, memo)
}

func dfs(i int, mask int, isLimit bool, isNum bool, s string, memo [][]int) int {

	if i == len(memo) {

		if isNum {
			return 1
		}
		return 0
	}

	if !isLimit && isNum && memo[i][mask] != -1 {
		return memo[i][mask]
	}
	res := 0
	if !isNum {
		res = dfs(i+1, mask, false, false, s, memo)
	}
	up := 9
	start := 0
	if isLimit {
		up = int(s[i] - '0')
	}
	if !isNum {
		start = 1
	}
	for j := start; j <= up; j++ {
		if mask&(1<<j) == 0 {
			res += dfs(i+1, mask|1<<j, isLimit && j == up, true, s, memo)
		}
	}

	if !isLimit && isNum {
		memo[i][mask] = res
	}
	return res
}
