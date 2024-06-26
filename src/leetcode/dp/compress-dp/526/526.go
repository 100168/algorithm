package main

import "math/bits"

/*
假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：

perm[i] 能够被 i 整除
i 能够被 perm[i] 整除
给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
*/
func countArrangement(n int) int {

	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, mask int) int {
		if i == 0 {

			//为什么需要特判，应为到这不一定满足情况
			if mask == (1<<n - 1) {
				return 1
			}
			return 0
		}

		if memo[i][mask] != -1 {
			return memo[i][mask]
		}
		cur := 0
		for j := 0; j < n; j++ {
			//注意这
			if 1<<j&mask == 0 && (j+1)%i == 0 || i%(j+1) == 0 {
				cur += dfs(i-1, mask|(1<<j))
			}
		}
		memo[i][mask] = cur
		return cur

	}
	return dfs(n, 0)
}

/*
*
不需要下标做法
*/
func countArrangement2(n int) int {

	memo := make([]int, 1<<n)

	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int

	dfs = func(mask int) int {

		if mask == 1<<n-1 {
			return 1
		}

		if memo[mask] != -1 {
			return memo[mask]
		}
		cur := 0
		index := bits.OnesCount(uint(mask)) + 1
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 && ((j+1)%index == 0 || index%(j+1) == 0) {
				cur += dfs(1<<j | mask)
			}
		}
		memo[mask] = cur
		return cur

	}

	return dfs(0)

}
func main() {
	countArrangement(2)
}
