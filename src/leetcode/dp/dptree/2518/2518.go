package main

func rootCount(edges [][]int, guesses [][]int, k int) int {

	n := len(edges) + 1
	g := make([][]int, n)

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
	}

	for _, v := range guesses {
		s[pair{v[0], v[1]}] = 1
	}

	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := make([]int, n)

	var dfs func(int, int)

	dfs = func(x, fa int) {

		for _, y := range g[x] {
			if y == fa {
				continue
			}
			if s[pair{x, y}] == 1 {
				cnt[0]++
			}
			dfs(y, x)
		}
	}

	ans := 0

	var reRoot func(int, int)

	reRoot = func(x, fa int) {

		for _, y := range g[x] {
			if y == fa {
				continue
			}
			cur := cnt[x]
			if s[pair{x, y}] == 1 {
				cur--
			}

			if s[pair{y, x}] == 1 {
				cur++
			}

			cnt[y] = cur

			if cnt[y] >= k {
				ans++
			}
			reRoot(y, x)

		}
	}

	dfs(0, -1)
	if cnt[0] >= k {
		ans++
	}
	reRoot(0, -1)
	return ans

}
