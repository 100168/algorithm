package main

/*
*
现有一棵 无向 树，树中包含 n 个节点，按从 0 到 n - 1 标记。树的根节点是节点 0 。给你一个长度为 n - 1 的二维整数数组 edges，其中 edges[i] = [ai, bi] 表示树中节点 ai 与节点 bi 之间存在一条边。

如果一个节点的所有子节点为根的
子树

	包含的节点数相同，则认为该节点是一个 好节点。

返回给定树中 好节点 的数量。

子树 指的是一个节点以及它所有后代节点构成的一棵树。
*/
func countGoodNodes(edges [][]int) int {

	n := len(edges) + 1

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := 0

	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		cnt := -1

		cc := 0
		flag := true
		for _, v := range g[x] {
			if v == fa {
				continue
			}
			child := dfs(v, x)

			if cnt == -1 {
				cnt = child
			} else {
				if cnt != child {
					flag = false
				}
			}
			cc += child

		}
		if flag {
			ans++
		}
		cc += 1
		return cc

	}

	dfs(0, -1)
	return ans

}
