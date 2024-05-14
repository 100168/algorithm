package main

/*
*
换根dp

给定一个无向、连通的树。树中有 n 个标记为 0...n-1 的节点以及 n-1 条边 。

给定整数 n 和数组 edges ， edges[i] = [ai, bi]表示树中的节点 ai 和 bi 之间有一条边。

返回长度为 n 的数组 answer ，其中 answer[i] 是树中第 i 个节点与所有其他节点之间的距离之和。
*/
func sumOfDistancesInTree(n int, edges [][]int) []int {

	ans := make([]int, n)
	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := make([]int, n)

	//第一次遍历计算根节点距离，和树大小
	var dfs func(int, int, int) int
	dfs = func(x, fa, deep int) int {
		cur := 1
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			cur += dfs(y, x, deep+1)
		}
		cnt[x] = cur
		ans[0] += deep
		return cur
	}

	//第二次遍历计算答案
	var findAsn func(int, int)

	findAsn = func(x, fa int) {
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			ans[y] = ans[x] + n - cnt[y]*2

			findAsn(y, x)
		}
	}

	dfs(0, -1, 0)
	findAsn(0, -1)
	return ans
}
