package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)

	up := make([]int, n)
	//向下延伸最大深度
	down1 := make([]int, n)
	//向下延伸第二大深度
	down2 := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	var dfs func(int, int) int

	dfs = func(x, fa int) int {

		first := 0
		second := 0
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			//当前最大最小
			f := dfs(y, x) + 1
			if f >= first {
				second = first
				first = f
			} else if f > second {
				second = f
			}
		}
		down1[x] = first
		down2[x] = second
		return first
	}

	dfs(0, -1)

	ans := make([]int, 0)
	ans = append(ans, 0)

	minVal := down1[0]
	var reRoot func(int, int)
	reRoot = func(x, fa int) {
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			//分类讨论  curX = max(up[x],down1[x])
			//由父节点的向上最大高度+1转变而来
			up[y] = up[x] + 1
			//根据x的最大深度是否跟y有关来更新up[y]的最大深度
			if down1[x]-1 == down1[y] {
				up[y] = max(up[y], down2[x]+1)
			} else {
				up[y] = max(up[y], down1[x]+1)
			}
			cur := max(up[y], down1[y])

			if cur < minVal {
				ans = make([]int, 0)
				minVal = cur
			}
			if minVal == cur {
				ans = append(ans, y)
			}
			reRoot(y, x)
		}
	}

	reRoot(0, -1)

	return ans

}
