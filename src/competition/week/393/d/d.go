package main

import "math"

func minimumValueSum(nums []int, andValues []int) int {
	n := len(nums)
	m := len(andValues)

	cache := make([][]map[int]int, n+1)
	for i := range cache {
		cache[i] = make([]map[int]int, m+1)
		for j := 0; j <= m; j++ {
			cache[i][j] = make(map[int]int)
		}
	}

	var dfs func(int, int, int) int

	dfs = func(i int, addi int, value int) int {

		if i >= n {
			if addi != m {
				return math.MaxInt / 2
			} else {
				return 0
			}
		}

		if _, ok := cache[i][addi][value]; ok {
			return cache[i][addi][value]
		}
		nd := value & nums[i]

		if addi >= m {
			return math.MaxInt / 2
		}

		if nd < andValues[addi] {
			return math.MaxInt / 2
		}

		cur := min(dfs(i+1, addi, nd))
		if nd == andValues[addi] {
			cur = min(cur, dfs(i+1, addi+1, -1)+nums[i])
		}
		cache[i][addi][value] = cur
		return cur
	}
	res := dfs(0, 0, -1)
	if res == math.MaxInt/2 {
		return -1
	}
	return res
}

func main() {
	println(minimumValueSum([]int{1, 4, 3, 3, 2}, []int{0, 3, 3, 2}))
}
