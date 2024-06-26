package main

import "fmt"

/*
*
有一棵由 n 个节点组成的无向树，以 0  为根节点，节点编号从 0 到 n - 1 。
给你一个长度为 n - 1 的二维 整数 数组 edges ，其中 edges[i] = [ai, bi]
表示在树上的节点 ai 和 bi 之间存在一条边。另给你一个下标从 0 开始、长度为 n 的数组 coins 和一个整数 k ，
其中 coins[i] 表示节点 i 处的金币数量。

从根节点开始，你必须收集所有金币。要想收集节点上的金币，必须先收集该节点的祖先节点上的金币。

节点 i 上的金币可以用下述方法之一进行收集：

收集所有金币，得到共计 coins[i] - k 点积分。如果 coins[i] - k 是负数，你将会失去 abs(coins[i] - k) 点积分。
收集所有金币，得到共计 floor(coins[i] / 2) 点积分。如果采用这种方法，
节点 i 子树中所有节点 j 的金币数 coins[j] 将会减少至 floor(coins[j] / 2) 。
返回收集 所有 树节点的金币之后可以获得的最大积分。
*/
func maximumPoints(edges [][]int, coins []int, k int) int {

	n := len(edges) + 1

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]

		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	memo := make([]map[int]int, n)

	for i := range memo {
		memo[i] = make(map[int]int)
	}
	var dfs func(int, int, int) int

	dfs = func(x, fa int, t int) int {

		if 1<<t > 10000 {
			return 0
		}
		if _, ok := memo[x][t]; ok {
			return memo[x][t]
		}

		coin := int(coins[x] / (1 << t))
		f0, f1 := coin-k, coin/2
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			f0 += dfs(y, x, t)
			f1 += dfs(y, x, t+1)
		}
		memo[x][t] = max(f0, f1)
		return max(f0, f1)
	}
	return dfs(0, -1, 0)
}

func main() {
	fmt.Println(maximumPoints([][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, []int{10, 10, 3, 20, 10}, 5))
}
