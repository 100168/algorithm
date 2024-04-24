package main

import "math"

/*
给你两个整数数组 nums1 和 nums2 ，它们长度都为 n 。

两个数组的 异或值之和 为 (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... + (nums1[n - 1] XOR nums2[n - 1]) （下标从 0 开始）。

比方说，[1,2,3] 和 [3,2,1] 的 异或值之和 等于 (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 + 2 = 4 。
请你将 nums2 中的元素重新排列，使得 异或值之和 最小 。

请你返回重新排列之后的 异或值之和 。
*/
func minimumXORSum(nums1 []int, nums2 []int) int {

	n := len(nums1)

	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, mark int) int {
		if i < 0 {
			return 0
		}

		if memo[i][mark] != -1 {
			return memo[i][mark]
		}

		cur := math.MaxInt / 2

		for j := 0; j < n; j++ {
			if 1<<j&mark == 0 {
				cur = min(cur, dfs(i-1, mark|(1<<j))+(nums1[i]^nums2[j]))
			}
		}

		memo[i][mark] = cur
		return cur
	}
	return dfs(n-1, 0)
}

func main() {
	println(minimumXORSum([]int{1, 2}, []int{2, 3}))
}
