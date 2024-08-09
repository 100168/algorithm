package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

/*
*
在一个小城市里，有 m 个房子排成一排，你需要给每个房子涂上 n 种颜色之一（颜色编号为 1 到 n ）。
有的房子去年夏天已经涂过颜色了，所以这些房子不可以被重新涂色。

我们将连续相同颜色尽可能多的房子称为一个街区。（比方说 houses = [1,2,2,3,3,2,1,1] ，它包含 5 个街区  [{1}, {2,2}, {3,3}, {2}, {1,1}] 。）

给你一个数组 houses ，一个 m * n 的矩阵 cost 和一个整数 target ，其中：

houses[i]：是第 i 个房子的颜色，0 表示这个房子还没有被涂色。
cost[i][j]：是将第 i 个房子涂成颜色 j+1 的花费。
请你返回房子涂色方案的最小总花费，使得每个房子都被涂色后，恰好组成 target 个街区。如果没有可用的涂色方案，请返回 -1 。

示例 1：

输入：houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
输出：9
解释：房子涂色方案为 [1,2,2,1,1]
此方案包含 target = 3 个街区，分别是 [{1}, {2,2}, {1,1}]。
涂色的总花费为 (1 + 1 + 1 + 1 + 5) = 9。
*/
func minCost(houses []int, cost [][]int, m int, n int, target int) int {

	memo := make([][][]int, m)

	for i := range memo {
		memo[i] = make([][]int, n+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, m)

			for x := range memo[i][j] {
				memo[i][j][x] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i int, preColor int, rest int) int {
		if rest < 0 {
			return math.MaxInt
		}
		if i < 0 {
			if rest == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		if memo[i][preColor][rest] != -1 {
			return memo[i][preColor][rest]
		}

		cur := math.MaxInt / 2
		for j := 0; j < n; j++ {
			if j != preColor {
				cur = min(cur, dfs(i-1, j, rest-1)+cost[i][j])
			} else {
				cur = min(cur, dfs(i-1, j, rest)+cost[i][j])
			}
		}
		if houses[i] != 0 {
			if j != preColor {
				cur = min(cur, dfs(i-1, j, rest-1)+cost[i][j])
			} else {
				cur = min(cur, dfs(i-1, j, rest)+cost[i][j])
			}
		}

		memo[i][preColor][rest] = cur
		return cur
	}
	return 1
}

func compute() {

	type pair struct {
		unitPrice float64
		qty       float64
	}
	posit := make([]pair, 10)
	amount1 := float64(0)
	qty1 := float64(0)
	unitPrice := float64(0)
	for i := range posit {
		price := float64(rand.Intn(100))
		qty := float64(rand.Intn(100))
		posit[i] = pair{price, qty}
		curAmount := amount1 + price*qty
		curQty := qty1 + qty
		curUnitPrice := curAmount / curQty
		amount1, qty1, unitPrice = curAmount, curQty, curUnitPrice
	}
	fmt.Printf("总价:%f,总数:%f,单价:%f", amount1, qty1, unitPrice)
	sort.Slice(posit, func(i, j int) bool {
		return posit[i].unitPrice < posit[i].unitPrice
	})

	amount1 = float64(0)
	qty1 = float64(0)
	unitPrice = float64(0)
	for i := range posit {
		price := posit[i].unitPrice
		qty := posit[i].qty
		posit[i] = pair{price, qty}
		curAmount := amount1 + price*qty
		curQty := qty1 + qty
		curUnitPrice := curAmount / curQty
		amount1, qty1, unitPrice = curAmount, curQty, curUnitPrice
	}
	fmt.Printf("总价:%f,总数:%f,单价:%f", amount1, qty1, unitPrice)
}

func main() {
	compute()
}
