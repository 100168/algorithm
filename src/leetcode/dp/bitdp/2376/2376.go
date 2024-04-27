package main

import (
	"math/bits"
	"strconv"
)

func countSpecialNumbers(n int) int {

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
			if mask>>j&1 == 0 {
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

func main() {
	println(bits.TrailingZeros(2))
}
