package main

import (
	"fmt"
	"sort"
)

/*
*
给你一个下标从 0 开始的二维整数数组 transactions，其中transactions[i] = [costi, cashbacki] 。

数组描述了若干笔交易。其中每笔交易必须以 某种顺序 恰好完成一次。在任意一个时刻，你有一定数目的钱 money ，为了完成交易 i ，money >= costi 这个条件必须为真。执行交易后，你的钱数 money 变成 money - costi + cashbacki 。

请你返回 任意一种 交易顺序下，你都能完成所有交易的最少钱数 money 是多少。

示例 1：
//-7,-14,-23,
// [8,1],[9,1],[7,1]
输入：transactions = [[2,1],[5,0],[4,2]]
输出：10
解释：
刚开始 money = 10 ，交易可以以任意顺序进行。
可以证明如果 money < 10 ，那么某些交易无法进行。

输入：transactions = [[3,0],[0,3]]
输出：3
解释：
- 如果交易执行的顺序是 [[3,0],[0,3]] ，完成所有交易需要的最少钱数是 3 。
- 如果交易执行的顺序是 [[0,3],[3,0]] ，完成所有交易需要的最少钱数是 0 。
所以，刚开始钱数为 3 ，任意顺序下交易都可以全部完成。

贪心:分两步
1.先将cost-cash<0 的分离开，然后按照cash从小到大排序
2.取所有cost-cash>=0中的cost最大值
*/
func minimumMoney(transactions [][]int) int64 {

	newT := make([][]int, 0)

	maC := 0
	for _, v := range transactions {
		if v[0] <= v[1] {
			maC = max(maC, v[0])
		} else {
			newT = append(newT, []int{v[0], v[1]})
		}
	}

	sort.Slice(newT, func(i, j int) bool {
		return newT[i][1] < newT[j][1]
	})

	ans := 0
	pre := 0
	for _, v := range newT {
		ans = min(ans, pre-v[0])
		pre += v[1] - v[0]
	}

	ans = min(ans, pre-maC)
	return -int64(ans)
}

// -3
func main() {
	fmt.Println(minimumMoney([][]int{{2, 1}, {5, 0}, {4, 2}}))
	fmt.Println(minimumMoney([][]int{{3, 0}, {0, 3}}))
}
