package main

import "fmt"

/*
*
Alice 和 Bob 再次设计了一款新的石子游戏。现有一行 n 个石子，每个石子都有一个关联的数字表示它的价值。
给你一个整数数组 stones ，其中 stones[i] 是第 i 个石子的价值。

Alice 和 Bob 轮流进行自己的回合，Alice 先手。每一回合，玩家需要从 stones 中移除任一石子。

如果玩家移除石子后，导致 所有已移除石子 的价值 总和 可以被 3 整除，那么该玩家就 输掉游戏 。
如果不满足上一条，且移除后没有任何剩余的石子，那么 Bob 将会直接获胜（即便是在 Alice 的回合）。
假设两位玩家均采用 最佳 决策。如果 Alice 获胜，返回 true ；如果 Bob 获胜，返回 false 。

stones = [2,1]
输出：true
解释：游戏进行如下：
- 回合 1：Alice 可以移除任意一个石子。
- 回合 2：Bob 移除剩下的石子。
已移除的石子的值总和为 1 + 2 = 3 且可以被 3 整除。因此，Bob 输，Alice 获胜。
*/
func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)

	type pair struct {
		x, y, z int
	}

	for _, v := range stones {
		cnt[v%3]++
	}

	if cnt[1] > 2 {
		if cnt[2] < cnt[1]-2 {
			s := cnt[2]*2 + cnt[0]
			if s%2 == 0 {
				return true
			}
		}
	}

	//112121212121212121212
	//221212121212121212121

}

func main() {
	fmt.Println(stoneGameIX([]int{2, 3}))
}
