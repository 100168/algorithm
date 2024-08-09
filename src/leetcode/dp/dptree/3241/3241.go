package main

/*
*
给你一棵 无向 树，树中节点从 0 到 n - 1 编号。同时给你一个长度为 n - 1 的二维整数数组 edges ，

	其中 edges[i] = [ui, vi] 表示节点 ui 和 vi 在树中有一条边。

一开始，所有 节点都 未标记 。对于节点 i ：

当 i 是奇数时，如果时刻 x - 1 该节点有 至少 一个相邻节点已经被标记了，那么节点 i 会在时刻 x 被标记。
当 i 是偶数时，如果时刻 x - 2 该节点有 至少 一个相邻节点已经被标记了，那么节点 i 会在时刻 x 被标记。
请你返回一个数组 times ，表示如果你在时刻 t = 0 标记节点 i ，那么时刻 times[i] 时，树中所有节点都会被标记。

请注意，每个 times[i] 的答案都是独立的，即当你标记节点 i 时，所有其他节点都未标记。

输入：edges = [[0,1],[0,2]]

输出：[2,4,3]

思路：这题目太难读了
*/
func timeTaken(edges [][]int) []int {

	n := len(edges) + 1
	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	type pair struct {
		x    int
		time int
	}
	first := make([]pair, n)

	second := make([]int, n)
	var dfs func(int, int) int

	dfs = func(x int, fa int) int {
		f := 0
		s := 0
		v := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			d := dfs(y, x) + 2 - y%2
			if d > f {
				s = f
				f = d
				v = y
			} else if d > s {
				s = d
			}
		}

		first[x] = pair{v, f}
		second[x] = s
		return f
	}

	ans := make([]int, n)
	var reRoot func(int, int, int)
	reRoot = func(x, fa, upFrom int) {
		p := first[x]
		ans[x] = max(p.time, upFrom)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			upMax := 0
			if first[x].x == y {
				upMax = second[x]
			} else {
				upMax = first[x].time
			}
			w := 2 - x%2
			reRoot(y, x, max(upFrom, upMax)+w)
		}
	}
	dfs(0, -1)
	reRoot(0, -1, 0)

	return ans
}

func timeTaken2(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	// nodes[x] 保存子树 x 的最大深度 maxD，次大深度 maxD2，以及最大深度要往儿子 y 走
	nodes := make([]struct{ maxD, maxD2, y int }, len(g))
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		p := &nodes[x]
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			maxD := dfs(y, x) + 2 - y%2 // 从 x 出发，往 y 方向的最大深度
			if maxD > p.maxD {
				p.maxD2 = p.maxD
				p.maxD = maxD
				p.y = y
			} else if maxD > p.maxD2 {
				p.maxD2 = maxD
			}
		}
		return p.maxD
	}
	dfs(0, -1)

	ans := make([]int, len(g))
	var reroot func(int, int, int)
	reroot = func(x, fa, fromUp int) {
		p := nodes[x]
		ans[x] = max(fromUp, p.maxD)
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			w := 2 - x%2  // 从 y 到 x 的边权
			if y == p.y { // 对于 y 来说，上面要选次大的
				reroot(y, x, max(fromUp, p.maxD2)+w)
			} else { // 对于 y 来说，上面要选最大的
				reroot(y, x, max(fromUp, p.maxD)+w)
			}
		}
	}
	reroot(0, -1, 0)
	return ans
}
