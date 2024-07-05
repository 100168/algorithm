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

思路：
1.贪心将a[i]+b[i] 从大到小排序

2.为什么这样选是最优？
3.假设只有两个数 a[i] b[i] a[j] b[j] 先选a[i] 则a-b = a[i] - b[i] 选a[j] 则 a - b a[j] - b[i]
  最优则：a[i]-b[j]>a[j]-b[i] ==>a[i]+b[i]>a[j]+b[j] ==>将a[i]+b[i] 从大到小排序


*/

func stoneGameVI(aliceValues []int, bobValues []int) int {

	n := len(aliceValues)
	a := 0
	b := 0
	sum := make([]int, n)

	for i := 0; i < n; i++ {
		b += bobValues[i]
		sum[i] = aliceValues[i] + bobValues[i]
	}
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	s := a - b

	for i := 0; i < n; i += 2 {
		s += sum[i]
	}

	if s < 0 {
		return -1
	}
	if s > 0 {
		return 1
	}
	return 0

}
