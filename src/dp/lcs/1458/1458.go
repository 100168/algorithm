package main

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {

	m := len(nums1)
	n := len(nums2)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, j int) int {

		if i == 0 && j == 0 {

			return nums1[i] * nums2[j]
		}
		if i < 0 || j < 0 {
			return math.MinInt / 2
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		cur := max(dfs(i-1, j-1)+nums1[i]*nums2[j], dfs(i-1, j), dfs(i, j-1))
		cache[i][j] = cur
		return cur
	}

	return dfs(m-1, n-1)
}

func main() {

	println(maxDotProduct([]int{-1, -1}, []int{1, 1}))
}
