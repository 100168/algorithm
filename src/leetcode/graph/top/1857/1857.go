package main

import (
	"fmt"
)

/*
*
给你一个 有向图 ，它含有 n 个节点和 m 条边。节点编号从 0 到 n - 1 。

给你一个字符串 colors ，其中 colors[i] 是小写英文字母，表示图中第 i 个节点的 颜色 （下标从 0 开始）。
同时给你一个二维数组 edges ，其中 edges[j] = [aj, bj] 表示从节点 aj 到节点 bj 有一条 有向边 。

图中一条有效 路径 是一个点序列 x1 -> x2 -> x3 -> ... -> xk ，
对于所有 1 <= i < k ，从 xi 到 xi+1 在图中有一条有向边。路径的 颜色值 是路径中 出现次数最多 颜色的节点数目。

请你返回给定图中有效路径里面的 最大颜色值 。如果图中含有环，请返回 -1 。
*/
func largestPathValue(colors string, edges [][]int) int {

	n := len(colors)
	g := make([][]int, n)

	inDegree := make([]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		inDegree[y]++
	}

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, 26)
	}
	var queue []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	ans := 1
	cnt := 0
	for len(queue) > 0 {
		cnt++
		cur := queue[0]
		queue = queue[1:]
		curColor := colors[cur] - 'a'
		dp[cur][curColor]++
		ans = max(ans, dp[cur][curColor])
		for _, v := range g[cur] {
			inDegree[v]--
			for i := 0; i < 26; i++ {
				dp[v][i] = max(dp[v][i], dp[cur][i])
			}
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if cnt != n {
		return -1
	}
	return ans
}

// 0->1
// 0->2->3->5->6->7->8
func main() {
	fmt.Println(largestPathValue("hhqhuqhqff", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {3, 5}, {5, 6}, {2, 7}, {6, 7}, {7, 8}, {3, 8}, {5, 8}, {8, 9}, {3, 9}, {6, 9}}))
}
