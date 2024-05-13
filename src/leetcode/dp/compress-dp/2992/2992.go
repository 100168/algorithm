package main

import "math/bits"

/*
*
给定一个整数 n，返回 下标从 1 开始 的数组 nums = [1, 2, ..., n]的 可能的排列组合数量，使其满足 自整除 条件。

如果对于每个 1 <= i <= n，满足 gcd(a[i], i) == 1，数组 nums 就是 自整除 的。

数组的 排列 是对数组元素的重新排列组合，例如，下面是数组 [1, 2, 3] 的所有排列组合
*/
func selfDivisiblePermutationCount(n int) int {

	memo := make([]int, 1<<n)

	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int

	dfs = func(mask int) int {
		index := bits.OnesCount(uint(mask))
		if index == n {
			return 1
		}

		if memo[mask] != -1 {
			return memo[mask]
		}
		cur := 0
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 && gcd(j+1, index+1) == 1 {

				cur += dfs(1<<j | mask)
			}
		}
		memo[mask] = cur
		return cur
	}
	return dfs(0)
}

func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
