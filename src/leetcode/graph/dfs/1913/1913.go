package main

/*
*
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，
其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆
，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。
*/
func makeConnected(n int, connections [][]int) int {

	edegs := len(connections)
	if edegs < n-1 {
		return -1
	}

	visited := make([]bool, n)

	g := make([][]int, n)
	for _, v := range connections {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int)

	dfs = func(x int) {

		visited[x] = true
		for _, y := range g[x] {
			if !visited[y] {
				dfs(y)
			}
		}
	}
	size := 0
	for i := 0; i < n; i++ {

		for !visited[i] {
			dfs(i)
			size++
		}
	}

	return size - 1
}
