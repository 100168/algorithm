package main

/*
给你一个整数 n ，表示一张 无向图 中有 n 个节点，
编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。

请你返回 无法互相到达 的不同 点对数目 。
*/

func countPairs(n int, edges [][]int) int64 {

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	visited := make([]bool, n)
	var dfs func(int) int

	dfs = func(x int) int {

		visited[x] = true
		size := 1

		for _, y := range g[x] {
			if !visited[y] {
				size += dfs(y)
			}

		}

		return size
	}

	total := 0
	ans := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			size := dfs(i)
			ans += total * size
			total += size

		}
	}
	return int64(ans)
}
func countPairs2(n int, edges [][]int) int64 {

	unitFind := new(UnitFind)
	unitFind.init(n)
	for _, edge := range edges {
		unitFind.unit(edge[0], edge[1])
	}

	cnt := int64(0)
	sizeList := unitFind.size
	sum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + sizeList[i-1]
	}
	for i := 0; i < n; i++ {
		cur := sizeList[i]

		cnt += int64(cur * (sum[n] - sum[i+1]))

	}
	return cnt
}

type UnitFind struct {
	parent []int
	size   []int
}

func (uinFind *UnitFind) init(n int) {
	uinFind.parent = make([]int, n)
	uinFind.size = make([]int, n)
	for i := 0; i < n; i++ {
		uinFind.parent[i] = i
		uinFind.size[i] = 1
	}
}

func (uinFind *UnitFind) find(a int) int {

	p := uinFind.parent
	for p[a] != a {
		p[a] = p[p[a]]
		a = p[a]
	}
	return a
}

func (uinFind *UnitFind) unit(a, b int) {

	findA := uinFind.find(a)
	findB := uinFind.find(b)
	if findA == findB {
		return
	}

	if uinFind.size[findA] < uinFind.size[findB] {
		uinFind.size[findB] += uinFind.size[findA]
		uinFind.size[findA] = 0
		uinFind.parent[findA] = findB
	} else {
		uinFind.size[findA] += uinFind.size[findB]
		uinFind.size[findB] = 0
		uinFind.parent[findB] = findA
	}

}
