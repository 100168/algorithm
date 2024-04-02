package main

import "sort"

func beautifulSubsets(nums []int, k int) int {

	sort.Ints(nums)
	n := len(nums)
	nMap := make(map[int]bool)
	cnt := 0
	for _, v := range nums {
		if nMap[v-k] {
			cnt++
		}
		nMap[v] = true
	}
	n -= cnt

	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	var dfs func(int) int

	dfs = func(i int) int {
		if i < 0 {
			return 1
		}

		if cache[i] != -1 {
			return cache[i]
		}
		cache[i] = 1 + dfs(i-1)
		return cache[i]
	}
	return dfs(n-1) * 1 << cnt
}

func main() {
	println(beautifulSubsets([]int{2, 4, 6, 8}, 2))
}
