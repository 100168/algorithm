package main

import (
	"fmt"
	"math"
)

/*
*
Alice 和 Bob 继续他们的石子游戏。几堆石子 排成一行 ，每堆石子都对应一个得分，由数组 stoneValue 给出。

Alice 和 Bob 轮流取石子，Alice 总是先开始。在每个玩家的回合中，该玩家可以拿走剩下石子中的的前 1、2 或 3 堆石子 。
比赛一直持续到所有石头都被拿走。

每个玩家的最终得分为他所拿到的每堆石子的对应得分之和。每个玩家的初始分数都是 0 。

比赛的目标是决出最高分，得分最高的选手将会赢得比赛，比赛也可能会出现平局。

假设 Alice 和 Bob 都采取 最优策略 。

如果 Alice 赢了就返回 "Alice" ，Bob 赢了就返回 "Bob"，分数相同返回 "Tie" 。
*/
func stoneGameIII(stoneValue []int) string {

	n := len(stoneValue)

	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = -1
	}

	var first func(int) int

	first = func(x int) int {
		if x == n {
			return 0
		}

		if f[x] != -1 {
			return f[x]
		}
		cur := math.MinInt

		pre := 0
		for i := x; i <= x+2 && i < n; i++ {
			pre += stoneValue[i]
			cur = max(cur, pre-first(i+1))
		}
		f[x] = cur
		return cur
	}

	res := first(0)

	if res == 0 {
		return "Tie"
	}
	if res > 0 {
		return "Alice"
	}
	return "Bob"
}

func main() {
	fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))
}
