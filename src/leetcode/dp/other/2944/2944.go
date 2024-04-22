package main

import (
	"fmt"
	"math"
)

/*
你在一个水果超市里，货架上摆满了玲琅满目的奇珍异果。

给你一个下标从 1 开始的数组 prices ，其中 prices[i] 表示你购买第 i 个水果需要花费的金币数目。

水果超市有如下促销活动：

如果你花费 price[i] 购买了水果 i ，那么后面的 i 个水果你可以免费获得。
注意 ，即使你 可以 免费获得水果 j ，你仍然可以花费 prices[j] 个金币去购买它以便能免费获得接下来的 j 个水果。

请你返回获得所有水果所需要的 最少 金币数。
*/
func minimumCoins(prices []int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt / 2
		}
	}
	//3,1,2,4,5,6,3,3,4,5,100,1,1,100,1001,100
	dp[1][1] = prices[0]
	for i := 2; i <= n; i++ {
		//i == 9           i-1==8   4                 5
		dp[i][1] = min(dp[i][1], dp[i-1][1]+prices[i-1], dp[i-1][0]+prices[i-1])
		for j := i - 1; j*2 >= i; j-- {
			if j*2 >= i {
				//不买
				dp[i][0] = min(dp[i][0], dp[j][1])
				//买
			}
		}
	}
	return min(dp[n][1], dp[n][0])
}

func minimumCoins2(prices []int) int {

	n := len(prices)

	memo := make([]int, (n+1)/2)

	for i := range memo {
		memo[i] = -1
	}
	//3,1,2,4,5,6,3,3,4,5,100,1,1,100,1001,100

	var dfs func(int) int

	dfs = func(i int) int {
		if i*2 >= n {
			return prices[i-1]
		}

		if memo[i] != -1 {
			return memo[i]
		}
		res := math.MaxInt
		//下一个需要购买的水果
		for j := i + 1; j <= i*2+1; j++ {
			res = min(res, dfs(j))
		}
		res += prices[i-1]

		memo[i] = res
		return res
	}
	return dfs(1)
}

func minimumCoins3(prices []int) int {
	n := len(prices)
	type pair struct{ i, f int }
	q := []pair{{n + 1, 0}} // 哨兵
	for i := n; i > 0; i-- {
		for q[0].i > i*2+1 { // 右边离开窗口
			q = q[1:]
		}
		f := prices[i-1] + q[0].f
		for f <= q[len(q)-1].f {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, f}) // 左边进入窗口
	}
	return q[len(q)-1].f
}

func main() {
	//fmt.Println(minimumCoins([]int{1, 2, 3}))
	//fmt.Println(minimumCoins([]int{1, 10, 1, 1}))
	fmt.Println(minimumCoins([]int{3, 1, 2, 4, 5, 7, 5}))
}
