package main

import (
	"math"
	"slices"
	"sort"
)

/*
924. 尽量减少恶意软件的传播
困难

相关标签

相关企业
给出了一个由 n 个节点组成的网络，用 n × n 个邻接矩阵图 graph 表示。在节点网络中，当 graph[i][j] = 1 时，表示节点 i 能够直接连接到另一个节点 j。

一些节点 initial 最初被恶意软件感染。只要两个节点直接连接，且其中至少一个节点受到恶意软件的感染，那么两个节点都将被恶意软件感染。这种恶意软件的传播将继续，直到没有更多的节点可以被这种方式感染。

假设 M(initial) 是在恶意软件停止传播之后，整个网络中感染恶意软件的最终节点数。

如果从 initial 中移除某一节点能够最小化 M(initial)， 返回该节点。如果有多个节点满足条件，就返回索引最小的节点。

请注意，如果某个节点已从受感染节点的列表 initial 中删除，它以后仍有可能因恶意软件传播而受到感染。



示例 1：

输入：graph = [[1,1,0],[1,1,0],[0,0,1]], initial = [0,1]
输出：0
示例 2：

输入：graph = [[1,0,0],[0,1,0],[0,0,1]], initial = [0,2]
输出：0
示例 3：

输入：graph = [[1,1,1],[1,1,1],[1,1,1]], initial = [1,2]
输出：1


提示：

n == graph.length
n == graph[i].length
2 <= n <= 300
graph[i][j] == 0 或 1.
graph[i][j] == graph[j][i]
graph[i][i] == 1
1 <= initial.length <= n
0 <= initial[i] <= n - 1
initial 中所有整数均不重复*/

func minMalwareSpread(graph [][]int, initial []int) int {

	n := len(graph)

	visited := make([]bool, n)

	initialMap := make(map[int]bool)
	for i := range initial {
		initialMap[initial[i]] = true
	}
	nodeId := 0
	size := 0
	var dfs func(int)

	dfs = func(x int) {

		size++
		visited[x] = true

		if nodeId != -2 && initialMap[x] {
			if nodeId < 0 {
				nodeId = x
			} else {
				nodeId = -2
			}
		}
		for y, v := range graph[x] {
			if v == 1 && !visited[y] {

				dfs(y)

			}
		}
	}
	ans := -1
	maxSize := 0
	for _, x := range initial {
		if visited[x] {
			continue
		}
		nodeId = -1
		size = 0
		dfs(x)
		if nodeId == x && (size > maxSize || size == maxSize && nodeId < ans) {
			ans = nodeId
			maxSize = size
		}
	}
	if ans < 0 {
		return slices.Min(initial)
	}

	return ans
}
func minMalwareSpread3(graph [][]int, initial []int) int {
	n := len(graph)
	u := unionFind{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	vis := make([]bool, n)
	for _, v := range initial {
		st := make([]int, 0)
		st = append(st, v)
		vis[v] = true
		for len(st) > 0 {
			temp := st
			st = nil
			for _, x := range temp {
				for y, f := range graph[x] {
					if vis[y] || f == 0 {
						continue
					}
					vis[y] = true
					u.union(x, y)
					st = append(st, y)
				}
			}
		}
	}
	sum := 0
	for _, v := range vis {
		if v {
			sum++
		}
	}
	ans := 0
	minV := math.MaxInt

	for i, x := range initial {
		flag := false
		for j, y := range initial {
			if j == i {
				continue
			}
			if u.isUnion(x, y) {
				flag = true
				break
			}
		}
		curIni := u.size[u.find(x)]
		if !flag {
			curIni = sum - curIni
		} else {
			curIni = sum
		}
		if curIni < minV {
			minV = curIni
			ans = x
		} else if curIni == minV {
			ans = min(ans, x)
		}
	}
	return ans
}

func minMalwareSpread2(graph [][]int, initial []int) int {
	n := len(graph)
	u := unionFind{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}
	for i := range graph {
		for j := range graph[i] {
			if graph[i][j] != 0 {
				u.union(i, j)
			}
		}
	}
	cnt := make([]int, n)
	sort.Ints(initial)
	ans := initial[0]
	minV := 0
	for _, x := range initial {
		cnt[u.find(x)]++
	}

	for _, x := range initial {
		cur := u.size[u.find(x)]
		if cnt[u.find(x)] == 1 && cur > minV {
			minV = cur
			ans = x
		}
	}
	return ans
}

type unionFind struct {
	parent []int
	size   []int
}

func (u unionFind) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u unionFind) isUnion(x, y int) bool {
	return u.find(x) == u.find(y)
}

func (u unionFind) union(a, b int) {
	findA := u.find(a)
	findB := u.find(b)
	if findA == findB {
		return
	}
	u.parent[findA] = findB
	u.size[findB] += u.size[findA]
	u.size[findA] = 1
}

func main() {
	println(minMalwareSpread2([][]int{{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		[]int{7, 8, 6, 2, 3}))
}
