package main

import "math/bits"

func maximumANDSum(nums []int, numSlots int) int {

	n := len(nums)

	numSlots *= 2
	memo := make([]int, 1<<numSlots)

	for i := range memo {
		memo[i] = -1

	}

	var dfs func(int) int

	dfs = func(mask int) int {
		i := bits.OnesCount(uint(mask))
		if i == n {
			return 0
		}
		if memo[mask] != -1 {
			return memo[mask]
		}

		cur := 0

		for j := 0; j < numSlots; j++ {
			if 1<<j&mask == 0 {
				cur = max(cur, dfs(mask|1<<j)+nums[i]&((j/2)+1))
			}
		}
		memo[mask] = cur
		return cur
	}
	return dfs(0)
}

func main() {
	println(maximumANDSum([]int{1, 3, 10, 4, 7, 1}, 9))
}
