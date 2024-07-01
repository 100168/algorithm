package main

/*
*
给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。

给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和  bi的人归入同一组。当可以用这种方法将所有人分进两组时，返回 true；否则返回 false。
*/
func possibleBiPartition(n int, dislikes [][]int) bool {

	colors := make([]int, n+1)

	var dfs func(int, int) bool

	g := make([][]int, n+1)
	for _, v := range dislikes {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	dfs = func(x, color int) bool {
		colors[x] = color
		for _, y := range g[x] {
			if colors[y] == 0 && !dfs(y, color^3) || colors[y] == color {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true

}
