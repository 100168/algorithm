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

思路：
1.构造题 想好怎么构造序列
2.然后1先还是2先分类讨论
*/
func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)
	for _, v := range stones {
		cnt[v%3]++
	}

	//1 12121212121212121212
	//2 21212121212121212121

	return check([]int{cnt[0], cnt[2], cnt[1]}) || check(cnt)

}

// todo
func stoneGameIX2(stones []int) bool {
	cnt := make([]int, 3)
	for _, v := range stones {
		cnt[v%3]++
	}

	if cnt[0]%2 == 0 {

	}

	return check([]int{cnt[0], cnt[2], cnt[1]}) || check(cnt)

}

func check(c []int) bool {
	if c[1] == 0 {
		return false
	}
	c[1]--                               // 开头为 1
	turn := 1 + min(c[1], c[2])*2 + c[0] // 计算回合数
	if c[1] > c[2] {                     // 序列末尾可以再加个 1
		turn++
		c[1]--
	}
	return turn%2 == 1 && c[1] != c[2] // 回合数为奇数，且还有剩余石子
}

func main() {
	fmt.Println(stoneGameIX([]int{15, 20, 10, 13, 14, 15, 5, 2, 3}))
}
