package main

import "strconv"

func countNumbersWithUniqueDigits(n int) int {

	n = pow(10, n) - 1
	s := strconv.Itoa(n)

	m := len(s)

	memo := make([][]int, m)

	for i := range memo {
		memo[i] = make([]int, 1<<10)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int

	dfs = func(i int, mask int, isLimit bool, isNum bool) int {

		if i == m {
			return 1
		}

		if !isLimit && isNum && memo[i][mask] != -1 {
			return memo[i][mask]
		}
		res := 0
		if !isNum {
			res += dfs(i+1, mask, false, false)
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		low := 0
		if !isNum {
			low = 1
		}

		for j := low; j <= up; j++ {

			if mask&(1<<j) == 0 {

				res += dfs(i+1, mask|1<<j, isLimit && j == up, true)
			}
		}
		if isNum && !isLimit {
			memo[i][mask] = res
		}
		return res
	}
	return dfs(0, 0, true, false)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 != 0 {
			res *= a
		}
		a *= a
		b = b >> 1
	}
	return res
}
