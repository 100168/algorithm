package main

var cnt []int

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {

	g := make([][]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	cnt = make([]int, n)
	for _, v := range trips {
		x, y := v[0], v[1]
		dfs(x, -1, y, g, price)
	}
	return min(minDfs(0, -1, g, price))
}

func dfs(x, fa, target int, g [][]int, price []int) bool {
	if x == target {
		cnt[x]++
		return true
	}

	for _, y := range g[x] {
		if y != fa && dfs(y, x, target, g, price) {
			cnt[x]++
			return true
		}
	}
	return false
}

// 跟打家劫舍3一样
func minDfs(x, fa int, g [][]int, price []int) (half, all int) {
	all = cnt[x] * price[x]
	half = all / 2

	for _, y := range g[x] {
		if y != fa {
			h, a := minDfs(y, x, g, price)
			half += a
			all += min(h, a)
		}
	}
	return half, all
}
