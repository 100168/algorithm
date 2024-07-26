package main

import (
	"fmt"
	"math"
)

/**
你需要制定一份 d 天的工作计划表。工作之间存在依赖，要想执行第 i 项工作，你必须完成全部 j 项工作（ 0 <= j < i）。

你每天 至少 需要完成一项任务。工作计划的总难度是这 d 天每一天的难度之和，而一天的工作难度是当天应该完成工作的最大难度。

给你一个整数数组 jobDifficulty 和一个整数 d，分别代表工作难度和需要计划的天数。第 i 项工作的难度是 jobDifficulty[i]。

返回整个工作计划的 最小难度 。如果无法制定工作计划，则返回 -1 。
*/

func minDifficulty(jobDifficulty []int, d int) int {

	n := len(jobDifficulty)

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, d+1)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i int, rest int) int {

		if i < 0 {
			if rest == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		if rest <= 0 {

			return math.MaxInt / 2
		}

		if f[i][rest] != -1 {
			return f[i][rest]
		}

		curMax := 0
		cur := math.MaxInt / 2

		for j := i; j >= 0; j-- {

			curMax = max(curMax, jobDifficulty[j])
			cur = min(cur, curMax+dfs(j-1, rest-1))
		}
		f[i][rest] = cur
		return cur
	}
	ans := dfs(n-1, d)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
}
