package main

import "sort"

func deleteAndEarn(nums []int) int {

	cntMap := make(map[int]int)

	newNums := make([]int, 0)

	for _, v := range nums {
		cntMap[v]++
	}
	for k := range cntMap {
		newNums = append(newNums, k)
	}
	n := len(newNums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	sort.Ints(newNums)

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		if cache[i] != -1 {
			return cache[i]
		}
		cur := 0
		j := i - 1
		if j >= 0 && newNums[j]+1 == newNums[i] {
			j--
		}
		cur = max(newNums[i]*cntMap[newNums[i]]+dfs(j), dfs(i-1))
		cache[i] = cur
		return cur
	}

	return dfs(n - 1)

}

func dp(nums []int) int {

	cntMap := make(map[int]int)

	newNums := make([]int, 0)

	for _, v := range nums {
		cntMap[v]++
	}
	for k := range cntMap {
		newNums = append(newNums, k)
	}
	n := len(newNums)
	sort.Ints(newNums)
	dp := make([]int, n)

	dp[0] = newNums[0] * cntMap[newNums[0]]
	for i := 1; i < n; i++ {
		j := i - 1
		if newNums[j]+1 == newNums[i] {
			j--
		}
		if j > 0 {
			dp[i] = newNums[i]*cntMap[newNums[i]] + dp[j]
		} else {
			dp[i] = newNums[i] * cntMap[newNums[i]]
		}
		dp[i] = max(dp[i], dp[i-1])
	}
	return dp[n-1]

}

func main() {
	println(deleteAndEarn([]int{1, 1, 1, 2, 4, 5, 5, 5, 6}))
	println(dp([]int{1, 1, 1, 2, 4, 5, 5, 5, 6}))
}
