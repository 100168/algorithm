package main

import (
	"strconv"
)

func findMaximumNumber(k int64, x int) int64 {

	l, r := 1, int(k<<x)

	for l <= r {
		mid := (r-l)/2 + l
		if check(mid, k, x) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return int64(r)
}

// mask   1...1...  =>...==k-1
func check(n int, k int64, x int) bool {

	s := strconv.FormatInt(int64(n), 2)

	n = len(s)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool, bool) int

	dfs = func(i int, sum int, isLimit bool, isNum bool) int {

		if i == n {
			return sum
		}

		if !isLimit && isNum && memo[i][sum] != -1 {
			return memo[i][sum]
		}

		res := 0

		if !isNum {
			res += dfs(i+1, sum, false, false)
		}
		low := 0
		up := 1

		if isLimit {
			up = int(s[i] - '0')
		}

		if !isNum {
			low = 1
		}

		for j := low; j <= up; j++ {

			if j == 1 && (n-i)%x == 0 {
				res += dfs(i+1, sum+1, isLimit && j == up, true)
			} else {
				res += dfs(i+1, sum, isLimit && j == up, true)
			}

		}
		if !isLimit && isNum {
			memo[i][sum] = res
		}
		return res
	}
	ans := dfs(0, 0, true, false)

	return int64(ans) <= k
}

func main() {
	println(findMaximumNumber(1, 5))
}
