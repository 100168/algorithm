package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {

	g := make([][]int, n)

	cnt := make([]int, n)
	for _, v := range edges {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
		cnt[y]++
	}
	ans := make([]map[int]bool, n)
	for i := range ans {
		ans[i] = make(map[int]bool)
	}
	st := make([]int, 0)
	for i, v := range cnt {
		if v == 0 {
			st = append(st, i)
		}
	}
	for len(st) > 0 {
		x := st[len(st)-1]
		st = st[:len(st)-1]
		for _, y := range g[x] {
			ans[y][x] = true
			for k := range ans[x] {
				ans[y][k] = true
			}
			cnt[y]--
			if cnt[y] == 0 {
				st = append(st, y)
			}
		}
	}
	res := make([][]int, n)
	for i, v := range ans {
		res[i] = make([]int, 0)
		for k := range v {
			res[i] = append(res[i], k)
		}
		sort.Ints(res[i])
	}
	return res
}

func main() {
	println(getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))
}
