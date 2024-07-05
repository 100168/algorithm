package main

/*
*
Alice 和 Bob 两个人轮流玩一个游戏，Alice 先手。

一开始，有 n 个石子堆在一起。每个人轮流操作，正在操作的玩家可以从石子堆里拿走 任意 非零 平方数 个石子。

如果石子堆里没有石子了，则无法操作的玩家输掉游戏。

给你正整数 n ，且已知两个人都采取最优策略。如果 Alice 会赢得比赛，那么返回 True ，否则返回 False 。

1 a  true
2 a false
3 a true 0*4+3
4 a true 1*4
5 a false 1*4+1
6 a true 1*4+2
7 a false  1*4+3
8 a true   2*4
9 a true   2*4+1
10 a false  2*4+2
11 a true   2*4+3
12 a false   3*4
13 a true  3*4 +1
14 a false 3*4+2
15 a true 3*4+3

狗贼，只能dp吗？
想了半天数学，我是废物
*/
func winnerSquareGame(n int) bool {

	f := make([]int, n+1)
	for i := range f {
		f[i] = -1
	}

	var dfs func(int) bool

	dfs = func(c int) bool {
		if c == 0 {
			return false
		}
		if c == 1 {
			return true
		}
		if f[c] != -1 {
			return f[c] == 1
		}
		cur := false
		for i := 1; i*i <= c; i++ {
			cur = cur || !dfs(c-i*i)
		}
		f[c] = 0
		if cur {
			f[c] = 1
		}
		return cur
	}

	return dfs(n)

}
