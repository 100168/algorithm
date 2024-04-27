package main

import (
	"fmt"
	"strconv"
)

func rotatedDigits(n int) int {

	s := strconv.Itoa(n)

	l := len(s)

	digits := []int{0, 1, 2, 5, 6, 8, 9}

	memo := make([][]int, l)

	for i := range memo {
		memo[i] = make([]int, 1<<8)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	invalid := 1<<0 | 1<<1 | 1<<5

	var dfs func(int, int, bool) int

	dfs = func(i, mask int, isLimit bool) int {
		if i == l {
			if mask|invalid == invalid {
				return 0
			}
			return 1
		}

		if !isLimit && memo[i][mask] != -1 {
			return memo[i][mask]
		}

		cur := 0

		up := 9

		if isLimit {
			up = int(s[i] - '0')
		}

		for j, d := range digits {
			if d > up {
				break
			}
			cur += dfs(i+1, mask|1<<j, isLimit && d == up)
		}

		if !isLimit {
			memo[i][mask] = cur
		}

		return cur

	}
	return dfs(0, 0, true)

}

func main() {

	fmt.Println(rotatedDigits(5))
}
