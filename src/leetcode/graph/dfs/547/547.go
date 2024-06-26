package main

/*
*
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，
且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，
其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。
*/
func findCircleNum(isConnected [][]int) int {

	n := len(isConnected)
	uf := new(unionFind)
	uf.parent = make([]int, n)
	uf.size = make([]int, n)

	for i := range uf.parent {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	for i := range isConnected {
		for j, v := range isConnected[i] {

			if v == 1 {
				uf.union(i, j)
			}
		}
	}
	ans := 0
	for _, v := range uf.size {
		if v > 0 {
			ans++
		}
	}
	return ans
}

type unionFind struct {
	parent []int
	size   []int
}

func (uf unionFind) find(x int) int {

	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf unionFind) union(a int, b int) {

	fa := uf.find(a)
	fb := uf.find(b)
	if fa == fb {
		return
	}

	if uf.size[fa] < uf.size[fb] {
		uf.parent[fa] = fb
		uf.size[fb] += uf.size[fa]
		uf.size[fa] = 0
	} else {
		uf.parent[fb] = fa
		uf.size[fa] += uf.size[fb]
		uf.size[fb] = 0
	}
}

func findCircleNum2(isConnected [][]int) int {

	n := len(isConnected)

	visited := make([]bool, n)

	var dfs func(x int)

	dfs = func(x int) {
		visited[x] = true
		for j, v := range isConnected[x] {
			if v == 1 && !visited[j] {
				dfs(j)
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			ans++
		}
	}
	return ans
}
