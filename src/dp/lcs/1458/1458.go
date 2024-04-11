package main

import (
	"slices"
)

func maxDotProduct(nums1 []int, nums2 []int) int {

	m := len(nums1)
	n := len(nums2)

	max1 := slices.Max(nums1)
	min1 := slices.Min(nums1)
	max2 := slices.Max(nums2)
	min2 := slices.Min(nums2)

	minVal := 0
	if max1 < 0 {
		minVal = max1 * min2
	}
	if max2 < 0 {
		minVal = min1 * max2
	}

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)

		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i, j int) int {

		if i < 0 || j < 0 {
			return 0
		}

		if cache[i][j] != -1 {
			return cache[i][j]
		}
		cur := max(dfs(i-1, j-1)+nums1[i]*nums2[j], dfs(i-1, j), dfs(i, j-1))

		cache[i][j] = cur
		return cur
	}

	ans := dfs(m-1, n-1)
	if ans == 0 {
		return minVal
	}
	return ans

}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	println(maxDotProduct([]int{-1, -1}, []int{1, 1}))
}
