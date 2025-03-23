package main

import "fmt"

/*
*
给你一个二维整数数组 properties，其维度为 n x m，以及一个整数 k。

定义一个函数 intersect(a, b)，它返回数组 a 和 b 中 共有的不同整数的数量 。

构造一个 无向图，其中每个索引 i 对应 properties[i]。如果且仅当 intersect(properties[i], properties[j]) >= k（其中 i 和 j 的范围为 [0, n - 1] 且 i != j），节点 i 和节点 j 之间有一条边。

返回结果图中 连通分量 的数量。
*/
func numberOfComponents(properties [][]int, k int) int {

	uf := new(uniFind)
	uf.init(len(properties))
	for i, v := range properties {
		for j := i + 1; j < len(properties); j++ {
			cnt := intersect(v, properties[j])
			if cnt >= k {
				uf.unit(i, j)
			}
		}
	}

	ans := 0
	for _, v := range uf.size {

		if v > 0 {
			ans++
		}
	}
	return ans

}

type uniFind struct {
	parent []int
	size   []int
}

func (uniFind *uniFind) init(n int) {
	uniFind.parent = make([]int, n)
	uniFind.size = make([]int, n)
	for i := range uniFind.parent {
		uniFind.parent[i] = i
		uniFind.size[i] = 1
	}
}

func (uniFind *uniFind) find(v int) int {
	parent := uniFind.parent
	for parent[v] != v {
		parent[v] = parent[parent[v]]
		v = parent[v]
	}
	return v
}

func (uniFind *uniFind) unit(a, b int) {
	findA := uniFind.find(a)
	findB := uniFind.find(b)
	if findA == findB {
		return
	}
	uniFind.parent[findA] = findB
	uniFind.size[findB] += uniFind.size[findA]
	uniFind.size[findA] = 0
}

func (uniFind *uniFind) isConnect(a, b int) bool {
	return uniFind.find(a) == uniFind.find(b)
}

func intersect(a, b []int) int {

	hs := make(map[int]bool)

	for _, v := range a {
		hs[v] = true
	}

	hb := make(map[int]bool)
	cnt := 0
	for _, v := range b {
		hb[v] = true
	}

	for k := range hb {
		if hs[k] {
			cnt++
		}
	}
	return cnt
}
func main() {
	fmt.Println(numberOfComponents([][]int{{1, 2}, {1, 1}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, 1))
}
