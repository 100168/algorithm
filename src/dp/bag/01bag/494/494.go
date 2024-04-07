package main

import "strconv"

func findTargetSumWays(nums []int, target int) int {

	n := len(nums)

	var dfs func(int, int) int

	cache := make(map[string]int)

	dfs = func(index, rest int) int {

		if index == n {
			if rest == 0 {
				return 1
			}
			return 0
		}

		key := strconv.Itoa(index) + "-" + strconv.Itoa(rest)
		if v, ok := cache[key]; ok {
			return v
		}

		cnt := 0
		cnt += dfs(index+1, rest-nums[index])
		cnt += dfs(index+1, rest+nums[index])
		cache[key] = cnt
		return cnt

	}
	return dfs(0, target)
}
func findTargetSumWays2(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	target += sum
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	var dfs func(int, int) int

	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, target+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	dfs = func(index, rest int) int {
		if rest < 0 {
			return 0
		}
		if index == n {
			if rest == 0 {
				return 1
			}
			return 0
		}
		if cache[index][rest] != -1 {
			return cache[index][rest]
		}

		cnt := 0
		cnt += dfs(index+1, rest)
		cnt += dfs(index+1, rest-nums[index])
		cache[index][rest] = cnt
		return cnt

	}
	return dfs(0, target)
}
