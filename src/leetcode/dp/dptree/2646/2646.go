package main

var cnt []int

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {

	g := make([][]int, n)

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := make([]int, n)

	var find func(int, int, int) bool

	find = func(x, fa, target int) bool {
		if x == target {
			cnt[x]++
			return true
		}
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			if find(y, x, target) {
				cnt[x]++
				return true
			}
		}
		return false
	}

	for _, v := range trips {
		find(v[0], -1, v[1])
	}

	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		noHalf := cnt[x] * price[x]
		half := noHalf / 2

		for _, y := range g[x] {
			if y == fa {
				continue
			}
			s1, s2 := dfs(y, x)
			half += s2
			noHalf += min(s1, s2)
		}
		return half, noHalf
	}
	return min(dfs(0, -1))
}
