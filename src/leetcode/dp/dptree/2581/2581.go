package main

/**

Alice 有一棵 n 个节点的树，节点编号为 0 到 n - 1 。树用一个长度为 n - 1 的二维整数数组 edges 表示，其中 edges[i] = [ai, bi] ，
表示树中节点 ai 和 bi 之间有一条边。
Alice 想要 Bob 找到这棵树的根。她允许 Bob 对这棵树进行若干次 猜测 。每一次猜测，Bob 做如下事情：
选择两个 不相等 的整数 u 和 v ，且树中必须存在边 [u, v] 。
Bob 猜测树中 u 是 v 的 父节点 。
Bob 的猜测用二维整数数组 guesses 表示，其中 guesses[j] = [uj, vj] 表示 Bob 猜 uj 是 vj 的父节点。
Alice 非常懒，她不想逐个回答 Bob 的猜测，只告诉 Bob 这些猜测里面 至少 有 k 个猜测的结果为 true 。
给你二维整数数组 edges ，Bob 的所有猜测和整数 k ，请你返回可能成为树根的 节点数目 。如果没有这样的树，则返回 0。
*/

func rootCount(edges [][]int, guesses [][]int, k int) int {

	n := len(edges) + 1
	g := make([][]int, n)

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses { // guesses 转成哈希表
		s[pair{p[0], p[1]}] = 1
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
