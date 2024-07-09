package main

/*
*
给你一棵 n 个节点的无向树，节点编号为 0 到 n - 1 。
给你整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 有一条边。

同时给你一个下标从 0 开始长度为 n 的整数数组 values ，其中 values[i] 是第 i 个节点的 值 。再给你一个整数 k 。

你可以从树中删除一些边，也可以一条边也不删，得到若干连通块。一个 连通块的值 定义为连通块中所有节点值之和。
如果所有连通块的值都可以被 k 整除，那么我们说这是一个 合法分割 。

请你返回所有合法分割中，连通块数目的最大值 。

我是废物
*/

func maxKDivisibleComponents3(n int, edges [][]int, values []int, k int) int {

	g := make([][]int, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	ans := 0
	var dfs func(int, int) int
	dfs = func(x, f int) int {
		cur := values[x]
		for _, v := range g[x] {
			if v != f {
				cur += dfs(v, x)
			}

		}
		if cur%k == 0 {
			ans++
		}
		return cur
	}
	dfs(0, -1)
	return ans

}
