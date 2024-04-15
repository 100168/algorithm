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
	del := make([]int, n)
	unDel := make([]int, n)

	unDel[0] = arr[0]

	maVal := slices.Max(arr)
	if maVal < 0 {
		return maVal
	}

	ans := math.MinInt
	for i := 1; i < n; i++ {
		unDel[i] = max(arr[i], unDel[i-1]+arr[i], del[i-1], 0)
		del[i] = max(del[i-1]+arr[i], arr[i], 0)
		ans = max(ans, del[i], unDel[i])
	}
	return ans
}
