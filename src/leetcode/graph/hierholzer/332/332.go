package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。
请你对该行程进行重新规划排序。

所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
如果存在多种有效的行程，请你按字典排序返回最小的行程组合。

例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
*/
func findItinerary(tickets [][]string) []string {

	g := make(map[string][]string)

	for _, v := range tickets {
		x, y := v[0], v[1]
		g[x] = append(g[x], y)
	}

	for s := range g {
		sort.Slice(g[s], func(i, j int) bool {
			return g[s][i] < g[s][j]
		})
	}

	var path []string
	var dfs func(x string)

	dfs = func(x string) {

		for len(g[x]) > 0 {
			next := g[x][0]
			g[x] = g[x][1:]
			dfs(next)
		}
		path = append(path, x)
	}
	dfs("JFK")
	slices.Reverse(path)
	return path
}

func main() {
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
}
