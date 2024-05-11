package main

import "fmt"

/*
*
有 n 个人排成一个队列，从左到右 编号为 0 到 n - 1 。
给你以一个整数数组 heights ，每个整数 互不相同，heights[i] 表示第 i 个人的高度。

一个人能 看到 他右边另一个人的条件是这两人之间的所有人都比他们两人 矮 。
更正式的，第 i 个人能看到第 j 个人的条件是 i < j 且 min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]) 。

请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是第 i 个人在他右侧队列中能 看到 的 人数 。
*/
func canSeePersonsCount(heights []int) []int {

	n := len(heights)
	ans := make([]int, n)
	st := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for len(st) > 0 && heights[st[len(st)-1]] <= heights[i] {
			cnt++
			st = st[:len(st)-1]
		}

		if len(st) > 0 {
			cnt++
		}
		ans[i] = cnt

		st = append(st, i)

	}
	return ans
}

func main() {

	fmt.Println(canSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
}
