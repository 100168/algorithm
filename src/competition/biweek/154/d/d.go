package main

import "fmt"

/*
*
给你一个整数 n 和一个以节点 1 为根的无向带权树，该树包含 n 个编号从 1 到 n 的节点。它由一个长度为 n - 1 的二维数组 edges 表示，
其中 edges[i] = [ui, vi, wi] 表示一条从节点 ui 到 vi 的无向边，权重为 wi。

同时给你一个二维整数数组 queries，长度为 q，其中每个 queries[i] 为以下两种之一：

[1, u, v, w'] – 更新 节点 u 和 v 之间边的权重为 w'，其中 (u, v) 保证是 edges 中存在的边。
[2, x] – 计算 从根节点 1 到节点 x 的 最短 路径距离。
返回一个整数数组 answer，其中 answer[i] 是对于第 i 个 [2, x] 查询，从节点 1 到 x 的最短路径距离。

示例 1：

输入： n = 2, edges = [[1,2,7]], queries = [[2,2],[1,1,2,4],[2,2]]

输出： [7,4]

1.假设一个月后你忘了这道题，你需要哪些关键点才能推导出解法

关键点：dfs时间戳、差分数组、数状数组
*/

func treeQueries(n int, edges [][]int, queries [][]int) []int {

	g := make([][]int, n+1)

	//建图
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)

	}
	//先序遍历，第一次来到该点的时间
	in := make([]int, n+1)
	//最后一次来到该点的时间
	//out -in = 该树的大小
	out := make([]int, n+1)

	time := 0
	var dfs func(int, int)
	dfs = func(x, f int) {
		time++
		in[x] = time
		for _, y := range g[x] {
			if y != f {
				dfs(y, x)
			}
		}
		out[x] = time
	}
	dfs(1, -1)

	bt := new(bitTree)
	bt.sum = make([]int, n+2)
	nums := make([]int, n+1)

	for _, v := range edges {

		x, y, d := v[0], v[1], v[2]

		if in[x] < in[y] {
			x = y
		}
		d -= nums[x]
		nums[x] = d
		bt.updateRange(in[x], out[x], d)

	}

	ans := make([]int, 0)

	for _, q := range queries {
		if q[0] == 1 {

			x, y, v := q[1], q[2], q[3]

			if in[x] < in[y] {
				x = y
			}
			l, r := in[x], out[x]
			bt.updateRange(l, r, v-nums[x])
			nums[x] = v
		} else {
			ans = append(ans, bt.query(in[q[1]]))
		}

	}
	return ans
}

type bitTree struct {
	sum []int
}

func (bt *bitTree) query(index int) int {

	x := 0
	for ; index > 0; index &= index - 1 {
		x += bt.sum[index]
	}
	return x
}
func (bt *bitTree) update(index, v int) {

	for ; index < len(bt.sum); index += index & -index {
		bt.sum[index] += v
	}
}

func (bt *bitTree) updateRange(l, r, v int) {

	bt.update(l, v)
	bt.update(r+1, -v)
}

func main() {
	fmt.Println(treeQueries(2, [][]int{{1, 2, 7}}, [][]int{{2, 2}, {1, 1, 2, 4}, {2, 2}}))
}
