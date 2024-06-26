package main

import (
	"fmt"
	"math"
)

/*
给你一个下标从 0 开始的二维整数数组 questions ，其中 questions[i] = [pointsi, brainpoweri] 。

这个数组表示一场考试里的一系列题目，你需要 按顺序 （也就是从问题 0 开始依次解决），针对每个问题选择 解决 或者 跳过 操作。解决问题 i 将让你 获得  pointsi 的分数，但是你将 无法 解决接下来的 brainpoweri 个问题（即只能跳过接下来的 brainpoweri 个问题）。如果你跳过问题 i ，你可以对下一个问题决定使用哪种操作。

比方说，给你 questions = [[3, 2], [4, 3], [4, 4], [2, 5]] ：
如果问题 0 被解决了， 那么你可以获得 3 分，但你不能解决问题 1 和 2 。
如果你跳过问题 0 ，且解决问题 1 ，你将获得 4 分但是不能解决问题 2 和 3 。
请你返回这场考试里你能获得的 最高 分数。
*/
func mostPoints(questions [][]int) int64 {
	n := len(questions)

	memo := make([]map[bool]int, n)

	for i := range memo {
		memo[i] = make(map[bool]int)
	}
	var dfs func(int, bool) int

	dfs = func(i int, flag bool) int {
		if i >= n {
			return 0
		}

		if _, ok := memo[i][flag]; ok {
			return memo[i][flag]
		}
		cur := 0
		end := questions[i][1]
		if flag {
			cur = max(dfs(i+end+1, true), dfs(i+end+1, false)) + questions[i][0]
		} else {
			cur = max(dfs(i+1, true), dfs(i+1, false))
		}
		memo[i][flag] = cur

		return cur
	}
	return int64(max(dfs(0, false), dfs(0, true)))
}

func main() {
	a := math.MaxInt32
	b := math.MaxInt32
	c := a + b
	fmt.Println(c)
}
