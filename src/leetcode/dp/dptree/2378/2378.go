package main

/*
给定一个 加权 树，由 n 个节点组成，从 0 到 n - 1。

该树以节点 0 为 根，用大小为 n 的二维数组 edges 表示，其中 edges[i] = [pari, weighti] 表示节点 pari 是节点 i 的 父 节点，
它们之间的边的权重等于 weighti。因为根结点 没有 父结点，所以有 edges[0] = [-1, -1]。

从树中选择一些边，使所选的两条边都不 相邻，所选边的权值之 和 最大。

返回所选边的 最大 和。

注意:

你可以 不选择 树中的任何边，在这种情况下权值和将为 0。
如果树中的两条边 Edge1 和 Edge2 有一个 公共 节点，它们就是 相邻 的。
换句话说，如果 Edge1连接节点 a 和 b, Edge2 连接节点 b 和 c，它们是相邻的。
*/
func maxScore(edges [][]int) int64 {

	n := len(edges)

	type pair struct{ y, d int }

	g := make([][]pair, n)

	for i, v := range edges {

		if i == 0 {
			continue
		}
		g[i] = append(g[i], pair{v[0], v[1]})
		g[v[0]] = append(g[v[0]], pair{i, v[1]})
	}

	var dfs func(int, int) (int, int)

	//1.选当前节点作为一条边 == 当前子节点不选 + 其他子节点的边选或不选取最大值并累加sum(max(选,不选))
	//2.不选当前节点作为一条边 == 所有子节点选或不选取最大值并累加sum(max(选，不选))
	dfs = func(x, fa int) (int, int) {

		//不选当前节点作为边
		no := 0
		//选当前节点作为边
		maxS := 0
		//子节点不选
		unMap := make(map[int]int)
		//子节点选
		sMap := make(map[int]int)
		//选当前节点作为边
		ssMap := make(map[int]int)
		for _, v := range g[x] {
			if v.y == fa {
				continue
			}
			y, n := dfs(v.y, x)
			sMap[v.y] = y
			unMap[v.y] = n
			ssMap[v.y] = n + v.d
			no += max(y, n)
		}
		//遍历一遍找出选哪个边值最大
		for k, v := range ssMap {
			maxS = max(maxS, v+no-max(unMap[k], sMap[k]))
		}
		return maxS, no
	}
	return int64(max(dfs(0, -1)))
}

func main() {
	println(maxScore([][]int{{-1, -1}, {0, 5}, {0, 10}, {2, 6}, {2, 4}}))
}
