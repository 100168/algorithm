package main

import "math"

/*
我们正在玩一个猜数游戏，游戏规则如下：

我从 1 到 n 之间选择一个数字。
你来猜我选了哪个数字。
如果你猜到正确的数字，就会 赢得游戏 。
如果你猜错了，那么我会告诉你，我选的数字比你的 更大或者更小 ，并且你需要继续猜数。
每当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。如果你花光了钱，就会 输掉游戏 。
给你一个特定的数字 n ，返回能够 确保你获胜 的最小现金数，不管我选择那个数字 。
*/
func getMoneyAmount(n int) int {

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i >= j {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		cur := math.MaxInt
		for k := i + 1; k < j; k++ {
			cur = min(max(dfs(i, k-1), dfs(k+1, j))+k, cur)
		}
		if i+1 == j {
			cur = i
		}
		memo[i][j] = cur
		return cur
	}
	return dfs(1, n)
}

func main() {
	println(getMoneyAmount(2))
}
