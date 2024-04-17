package main

import (
	"slices"
)

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
