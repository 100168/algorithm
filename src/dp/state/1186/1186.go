package main

import (
	"math"
	"slices"
)

func maximumSum(arr []int) int {

	n := len(arr)
	cache := make([]map[bool]int, n)
	for i := range cache {
		cache[i] = make(map[bool]int)
	}
	maxV := slices.Max(arr)
	if maxV < 0 {
		return maxV
	}
	var dfs func(int, bool) int
	dfs = func(i int, flag bool) int {
		if i < 0 {
			return 0
		}
		if _, ok := cache[i][flag]; ok {
			return cache[i][flag]
		}
		cur := math.MinInt
		cur = max(dfs(i-1, flag)+arr[i], arr[i], 0)
		if !flag {
			cur = max(dfs(i-1, true), cur)
		}
		cache[i][flag] = cur
		return cur
	}

	ans := math.MinInt

	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i, false))
	}
	return ans
}

func maximumSum2(arr []int) int {

	n := len(arr)
	del := make([]int, n+1)
	unDel := make([]int, n+1)
	unDel[0] = math.MinInt / 2
	del[0] = math.MinInt / 2
	ans := math.MinInt
	for i := 0; i < n; i++ {
		unDel[i+1] = max(unDel[i], 0) + arr[i]
		del[i+1] = max(del[i]+arr[i], unDel[i])
		ans = max(ans, del[i+1], unDel[i+1])
	}
	return ans
}
