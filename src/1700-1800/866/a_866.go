package main

/*给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。

给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和  bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。



示例 1：

输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
输出：true
解释：group1 [1,4], group2 [2,3]
示例 2：

输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
输出：false
示例 3：

输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
输出：false

1. 染色法
2. 并查集
*/

var colors []int

func possibleBiPartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	colors = make([]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	for i, color := range colors {
		if color == 0 && !dfs(i, g, 1) {
			return false
		}
	}

	return true
}

func dfs(parent int, next [][]int, color int) bool {
	for _, v := range next[parent] {
		if colors[v] == 0 && !dfs(v, next, 3^color) || colors[v] == color {
			return false
		}
	}
	return true
}

func possibleBiPartition2(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	uniFind := new(uniFind)
	uniFind.init(n)
	for i, v := range g {
		for _, next := range v {
			if uniFind.isConnect(i, next) {
				return false
			}
			uniFind.unit(next, v[0])
		}
	}
	return true
}

type uniFind struct {
	parent []int
}

func (uniFind *uniFind) init(n int) {
	uniFind.parent = make([]int, n+1)
	for i := range uniFind.parent {
		uniFind.parent[i] = i
	}
}

func (uniFind *uniFind) find(v int) int {
	parent := uniFind.parent
	for parent[v] != v {
		parent[v] = parent[parent[v]]
		v = parent[v]
	}
	return v
}

func (uniFind *uniFind) unit(a, b int) {
	findA := uniFind.find(a)
	findB := uniFind.find(b)
	if findA == findB {
		return
	}
	uniFind.parent[findA] = findB
}

func (uniFind *uniFind) isConnect(a, b int) bool {
	return uniFind.find(a) == uniFind.find(b)
}
