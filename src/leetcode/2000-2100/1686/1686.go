package main

import "sort"

/**
Alice 和 Bob 轮流玩一个游戏，Alice 先手。

一堆石子里总共有 n 个石子，轮到某个玩家时，他可以 移出 一个石子并得到这个石子的价值。
Alice 和 Bob 对石子价值有 不一样的的评判标准 。双方都知道对方的评判标准。

给你两个长度为 n 的整数数组 aliceValues 和 bobValues 。aliceValues[i] 和 bobValues[i] 分别表示 Alice 和 Bob 认为第 i 个石子的价值。

所有石子都被取完后，得分较高的人为胜者。如果两个玩家得分相同，那么为平局。两位玩家都会采用 最优策略 进行游戏。

请你推断游戏的结果，用如下的方式表示：

如果 Alice 赢，返回 1 。
如果 Bob 赢，返回 -1 。
如果游戏平局，返回 0 。

输入：aliceValues = [2,4,3], bobValues = [1,6,7]

[1,-2,-3]

[1,2]
[3,1]


输出：-1
解释：
不管 Alice 怎么操作，Bob 都可以得到比 Alice 更高的得分。
比方说，Alice 拿石子 1 ，Bob 拿石子 2 ， Alice 拿石子 0 ，Alice 会得到 6 分而 Bob 得分为 7 分。
Bob 会获胜。
*/

func stoneGameVI(aliceValues []int, bobValues []int) int {

	n := len(aliceValues)
	s1 := 0
	s2 := 0

	type pair struct {
		alice, bob int
	}
	pairs := make([]pair, n)

	for i := 0; i < n; i++ {
		s1 += aliceValues[i]
		s2 += bobValues[i]
		pairs[i] = pair{aliceValues[i], bobValues[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].alice-pairs[i].bob > pairs[j].alice-pairs[j].bob
	})

	l, r := 0, n-1

	for l < r {
		alice := pairs[l]
		bob := pairs[r]
		s2 -= alice.bob
		s1 -= bob.alice
		l++
		r--

	}

	if l == r {
		alice := pairs[l]
		s2 -= alice.bob
	}

	if s1 < s2 {
		return -1
	}
	if s1 == s2 {
		return 0
	}
	return 1

}
