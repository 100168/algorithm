package main

import "math"

/*
*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。
*/
func numSquares(n int) int {

	memo := make([]int, n+1)

	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int

	dfs = func(n int) int {
		if n == 0 {
			return 0
		}

		if memo[n] != -1 {
			return memo[n]
		}
		cur := math.MaxInt
		for i := 1; i*i <= n; i++ {
			cur = min(cur, dfs(n-i*i)+1)
		}
		memo[n] = cur
		return cur
	}
	return dfs(n)
}

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquares2(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
