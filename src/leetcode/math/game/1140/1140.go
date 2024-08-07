package main

import (
	"fmt"
	"math"
)

/*
*
爱丽丝和鲍勃继续他们的石子游戏。许多堆石子 排成一行，每堆都有正整数颗石子 piles[i]。游戏以谁手中的石子最多来决出胜负。

爱丽丝和鲍勃轮流进行，爱丽丝先开始。最初，M = 1。

在每个玩家的回合中，该玩家可以拿走剩下的 前 X 堆的所有石子，其中 1 <= X <= 2M。然后，令 M = max(M, X)。

游戏一直持续到所有石子都被拿走。

假设爱丽丝和鲍勃都发挥出最佳水平，返回爱丽丝可以得到的最大数量的石头。

输入：piles = [2,7,9,4,4]
输出：10
解释：如果一开始Alice取了一堆，Bob取了两堆，然后Alice再取两堆。
爱丽丝可以得到2 + 4 + 4 = 10堆。如果Alice一开始拿走了两堆，那么Bob可以拿走剩下的三堆。
在这种情况下，Alice得到2 + 7 = 9堆。返回10，因为它更大。
*/
func stoneGameII(piles []int) int {

	n := len(piles)

	f := make([][]int, n)

	for i := range f {
		f[i] = make([]int, n+1)

		for j := range f[i] {
			f[i][j] = -1
		}
	}

	s := 0

	for _, v := range piles {
		s += v
	}
	var dfs func(i, m int) int
	dfs = func(i, m int) int {

		if i == n {
			return 0
		}

		if f[i][m] != -1 {
			return f[i][m]
		}
		cur := math.MinInt
		pre := 0
		for j := i; j < i+2*m && j < n; j++ {
			pre += piles[j]
			cur = max(pre-dfs(j+1, max(j-i+1, m)), cur)
		}
		f[i][m] = cur
		return cur
	}
	//x+y==s x-y==d

	return (s + dfs(0, 1)) / 2
}

func main() {
	fmt.Println(stoneGameII([]int{1}))
}
