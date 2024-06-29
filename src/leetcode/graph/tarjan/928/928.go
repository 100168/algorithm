package main

import (
	"slices"
)

/*
*给定一个由 n 个节点组成的网络，用 n x n 个邻接矩阵 graph 表示。在节点网络中，
只有当 graph[i][j] = 1 时，节点 i 能够直接连接到另一个节点 j。

一些节点 initial 最初被恶意软件感染。只要两个节点直接连接，且其中至少一个节点受到恶意软件的感染，
那么两个节点都将被恶意软件感染。这种恶意软件的传播将继续，直到没有更多的节点可以被这种方式感染。

假设 M(initial) 是在恶意软件停止传播之后，整个网络中感染恶意软件的最终节点数。

我们可以从 initial 中 删除一个节点，并完全移除该节点以及从该节点到任何其他节点的任何连接。

请返回移除后能够使 M(initial) 最小化的节点。如果有多个节点满足条件，返回索引 最小的节点 。
*/
func minMalwareSpread(graph [][]int, initial []int) int {
	n := len(graph)
	vis := make([]bool, n)
	isInitial := make([]bool, n)
	for _, x := range initial {
		isInitial[x] = true
	}

	var nodeId, size int
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		size++
		for y, conn := range graph[x] {
			if conn == 0 {
				continue
			}
			if isInitial[y] {
				// 按照 924 题的状态机更新 nodeId
				// 注意避免重复统计，例如上图中的 0 有两条不同路径可以遇到 1
				if nodeId != -2 && nodeId != y {
					if nodeId == -1 {
						nodeId = y
					} else {
						nodeId = -2
					}
				}
			} else if !vis[y] {
				dfs(y)
			}
		}
	}
	cnt := make([]int, n)
	for i, seen := range vis {
		if seen || isInitial[i] {
			continue
		}
		//妙啊
		nodeId = -1
		size = 0
		dfs(i)
		if nodeId >= 0 { // 只找到一个在 initial 中的节点
			// 删除节点 nodeId 可以让 size 个点不被感染
			cnt[nodeId] += size
		}
	}

	maxCnt := 0
	minNodeId := -1
	for i, c := range cnt {
		if c > maxCnt {
			maxCnt = cnt[i]
			minNodeId = i
		}
	}
	if minNodeId >= 0 {
		return minNodeId
	}
	return slices.Min(initial)
}
