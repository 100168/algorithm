package main

/**


给你一棵 无向 树，树中节点从 0 到 n - 1 编号。同时给你一个长度为 n - 1 的二维整数数组 edges ，
其中 edges[i] = [ui, vi] 表示节点 ui 和 vi 在树中有一条边。

一开始，所有 节点都 未标记 。对于节点 i ：

当 i 是奇数时，如果时刻 x - 1 该节点有 至少 一个相邻节点已经被标记了，那么节点 i 会在时刻 x 被标记。
当 i 是偶数时，如果时刻 x - 2 该节点有 至少 一个相邻节点已经被标记了，那么节点 i 会在时刻 x 被标记。
请你返回一个数组 times ，表示如果你在时刻 t = 0 标记节点 i ，那么时刻 times[i] 时，树中所有节点都会被标记。

请注意，每个 times[i] 的答案都是独立的，即当你标记节点 i 时，所有其他节点都未标记。
*/

func timeTaken(edges [][]int) []int {

	n := len(edges) + 1

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	//vis := make([]bool, n)
	//
	//var find func(int, int, int)
	//
	//find := func(x, fa, t int) {
	//
	//	if x%2 == 0 {
	//
	//	}
	//	for _, v := range g[x] {
	//
	//		if v == fa {
	//			continue
	//		}
	//
	//	}
	//}
	return []int{0, 0}

}
