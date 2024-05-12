package main

import (
	"fmt"
	"math"
)

func findPermutation(nums []int) []int {

	n := len(nums)
	memo := make([][]int, 1<<n)
	for i := range memo {
		memo[i] = make([]int, n)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(mask int, pre int) int {

		if mask == 1<<n-1 {
			return abs(pre - nums[0])
		}

		if memo[mask][pre] != -1 {
			return memo[mask][pre]
		}

		cur := math.MaxInt
		for j := 0; j < n; j++ {
			if 1<<j&mask == 0 {
				cur = min(cur, dfs(1<<j|mask, j)+abs(pre-nums[j]))
			}
		}
		memo[mask][pre] = cur
		return cur
	}

	var ans []int
	var build func(int, int)

	build = func(mask int, pre int) {
		ans = append(ans, pre)
		if mask == 1<<n-1 {
			return
		}
		k := dfs(mask, pre)
		for j := 0; j < n; j++ {
			if 1<<j&mask != 0 {
				continue
			}
			if dfs(1<<j|mask, j)+abs(pre-nums[j]) == k {
				build(1<<j|mask, j)
				break
			}
		}

	}

	dfs(1, 0)
	build(1, 0)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	fmt.Println(findPermutation([]int{1, 0, 2}))

}
