package main

import (
	"fmt"
)

/*
n 块石头放置在二维平面中的一些整数坐标点上。每个坐标点上最多只能有一块石头。

如果一块石头的 同行或者同列 上有其他石头存在，那么就可以移除这块石头。

给你一个长度为 n 的数组 stones ，其中 stones[i] = [xi, yi] 表示第 i 块石头的位置，返回 可以移除的石子 的最大数量。


输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
输出：5
解释：一种移除 5 块石头的方法如下所示：
1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
石头 [0,0] 不能移除，因为它没有与另一块石头同行/列。
*/

func removeStones(stones [][]int) int {

	rowMap := make(map[int][]int)
	colMap := make(map[int][]int)
	ans := 0

	n := len(stones)
	uf := new(uniFind)
	uf.init(n)
	for i, v := range stones {
		r, c := v[0], v[1]
		for _, v := range rowMap[r] {
			uf.unit(i, v)
		}
		for _, v := range colMap[r] {
			uf.unit(i, v)
		}
		rowMap[r] = append(rowMap[r], i)
		colMap[c] = append(colMap[c], i)
	}

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
	uniFind.parent = make([]int, n+1)
	uniFind.size = make([]int, n+1)
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
	uniFind.size[findA] = 0
}

func main() {
	//fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
	fmt.Println(removeStones([][]int{{0, 1}, {1, 0}, {1, 1}}))
}
