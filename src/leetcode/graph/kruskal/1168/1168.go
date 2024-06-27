package main

import (
	"sort"
)

/*
*
村里面一共有 n 栋房子。我们希望通过建造水井和铺设管道来为所有房子供水。
对于每个房子 i，我们有两种可选的供水方案：一种是直接在房子内建造水井，
成本为 wells[i - 1] （注意 -1 ，因为 索引从0开始 ）；
另一种是从另一口井铺设管道引水，数组 pipes 给出了在房子间铺设管道的成本，
其中每个 pipes[j] = [house1j, house2j, costj] 代表用管道将 house1j 和 house2j连接在一起的成本。连接是双向的。

请返回 为所有房子都供水的最低总成本 。
*/
func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {

	e := make([]pair, n+1)
	for _, v := range pipes {
		x, y, d := v[0]-1, v[1]-1, v[2]
		e = append(e, pair{x, y, d})
	}

	for i, d := range wells {
		e = append(e, pair{i, n, d})
	}
	uf := new(unionFind)
	uf.parent = make([]int, n+1)
	for i := range uf.parent {
		uf.parent[i] = i
	}

	sort.Slice(e, func(i, j int) bool {
		return e[i].d < e[j].d
	})
	ans := 0
	for _, v := range e {
		if uf.find(v.x) != uf.find(v.y) {
			uf.union(v.x, v.y)
			ans += v.d
		}
	}
	return ans
}

type pair struct {
	x int
	y int
	d int
}

type unionFind struct {
	parent []int
}

func (uf unionFind) find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}
func (uf unionFind) union(a, b int) bool {
	fa := uf.find(a)
	fb := uf.find(b)
	if fa == fb {
		return false
	}
	uf.parent[fa] = fb
	return true
}
