package main

import "math"

func maxValueOfCoins(piles [][]int, k int) int {

	n := len(piles)
	var dfs func(int, int) int

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	dfs = func(i int, rest int) int {
		if rest == 0 {
			return 0
		}

		if i < 0 {
			return math.MinInt
		}

		if cache[i][rest] != -1 {
			return cache[i][rest]
		}
		cur := 0
		pile := piles[i]

		preSum := 0
		m := len(pile)
		cur = max(cur, dfs(i-1, rest))
		for j := 1; j <= min(m, rest); j++ {
			preSum += pile[j-1]
			cur = max(dfs(i-1, rest-j)+preSum, cur)
		}
		cache[i][rest] = cur
		return cur
	}
	return dfs(n-1, k)
}

func maxValueOfCoins2(piles [][]int, k int) int {

	dp := make([]int, k+1)
	for i := range dp {
		dp[i] = math.MinInt
	}
	dp[0] = 0
	for _, pile := range piles {
		for x := k; x >= 0; x-- {
			sum := 0
			for j := 1; j <= min(len(pile), x); j++ {
				sum += pile[j-1]
				dp[x] = max(dp[x], dp[x-j]+sum)
			}
		}

	}
	return dp[k]
}

func main() {
	println(maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
	println(maxValueOfCoins2([][]int{{1, 100, 3}, {7, 8, 9}}, 2))
}
