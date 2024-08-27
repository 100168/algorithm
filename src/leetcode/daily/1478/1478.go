package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
给你一个房屋数组houses 和一个整数 k ，其中 houses[i] 是第 i 栋房子在一条街上的位置，现需要在这条街上安排 k 个邮筒。

请你返回每栋房子与离它最近的邮筒之间的距离的 最小 总和。

答案保证在 32 位有符号整数范围以内。'

1.思路dp+中位数贪心
*/
func minDistance(houses []int, k int) int {

	n := len(houses)

	sort.Ints(houses)

	sum := make([][]int, n)

	for i := range sum {
		sum[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			sum[i][j] = sum[i+1][j-1] + houses[j] - houses[i]
		}
	}

	f := make([][]int, n)
	for i := range f {

		f[i] = make([]int, k)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int

	dfs = func(i int, j int) int {

		if i < 0 {
			return 0
		}

		if j < 0 {
			return math.MaxInt / 2
		}

		if f[i][j] != -1 {
			return f[i][j]
		}

		cur := math.MaxInt / 2

		for t := i; t >= 0; t-- {

			cur = min(cur, dfs(t-1, j-1)+sum[t][i])
		}
		f[i][j] = cur
		return cur
	}

	return dfs(n-1, k-1)

}
func main() {
	fmt.Println(minDistance([]int{1, 4, 8, 10, 20}, 3))
}
