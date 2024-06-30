package main

/*
*给你两棵 无向 树，分别有 n 和 m 个节点，节点编号分别为 0 到 n - 1 和 0 到 m - 1 。
给你两个二维整数数组 edges1 和 edges2 ，长度分别为 n - 1 和 m - 1 ，
其中 edges1[i] = [ai, bi] 表示在第一棵树中节点 ai 和 bi 之间有一条边，
edges2[i] = [ui, vi] 表示在第二棵树中节点 ui 和 vi 之间有一条边。

你必须在第一棵树和第二棵树中分别选一个节点，并用一条边连接它们。

请你返回添加边后得到的树中，最小直径 为多少。

一棵树的 直径 指的是树中任意两个节点之间的最长路径长度。
*/
func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	d1 := minLen(edges1)
	d2 := minLen(edges2)
	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}

func minLen(e [][]int) int {
	n := len(e) + 1
	g := make([][]int, n)
	for _, v := range e {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 0
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		maxLen := 0
		for _, y := range g[x] {
			if y != fa {
				next := dfs(y, x) + 1
				ans = max(ans, maxLen+next)
				maxLen = max(maxLen, next)
			}
		}
		return maxLen
	}
	dfs(0, -1)
	return ans
}
