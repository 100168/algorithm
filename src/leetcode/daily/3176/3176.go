package main

import "fmt"

func maximumLength(nums []int, k int) int {

	n := len(nums)

	memo := make([][][]int, n)

	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for m := range memo[i][j] {
				memo[i][j][m] = -1
			}
		}
	}
	var dfs func(int, int, int) int

	dfs = func(i, pre, rest int) int {

		if i < 0 {
			return 0
		}

		if pre >= 0 && memo[i][pre][rest] != -1 {
			return memo[i][pre][rest]
		}
		cur := 0

		cur = dfs(i-1, pre, rest)

		if pre == -1 || nums[i] == nums[pre] {
			cur = max(cur, dfs(i-1, i, rest)+1)
		} else {
			if rest-1 >= 0 {
				cur = max(cur, dfs(i-1, i, rest-1)+1)
			}

		}

		if pre >= 0 {
			memo[i][pre][rest] = cur
		}

		return cur
	}
	return dfs(n-1, -1, k)
}

func main() {
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 3}, 2))
}
