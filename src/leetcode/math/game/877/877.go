package main

/*
*
Alice 和 Bob 用几堆石子在做游戏。一共有偶数堆石子，排成一行；每堆都有 正 整数颗石子，数目为 piles[i] 。

游戏以谁手中的石子最多来决出胜负。石子的 总数 是 奇数 ，所以没有平局。

Alice 和 Bob 轮流进行，Alice 先开始 。 每回合，玩家从行的 开始 或 结束 处取走整堆石头。
这种情况一直持续到没有更多的石子堆为止，此时手中 石子最多 的玩家 获胜 。

假设 Alice 和 Bob 都发挥出最佳水平，当 Alice 赢得比赛时返回 true ，当 Bob 赢得比赛时返回 false 。

输入：piles = [5,3,4,5]
输出：true
解释：
Alice 先开始，只能拿前 5 颗或后 5 颗石子 。
假设他取了前 5 颗，这一行就变成了 [3,4,5] 。
如果 Bob 拿走前 3 颗，那么剩下的是 [4,5]，Alice 拿走后 5 颗赢得 10 分。
如果 Bob 拿走后 5 颗，那么剩下的是 [3,4]，Alice 拿走后 4 颗赢得 9 分。
这表明，取前 5 颗石子对 Alice 来说是一个胜利的举动，所以返回 true 。

3 5 4 1
先手必胜

思路：
1.将石头根据下标奇偶进行分组
2.a先手可以决定取哪堆   如果取奇数堆，则后手只能取偶数堆，如果取偶数堆后手只能取奇数堆， 所以先手可以决定一直选哪堆，先手必胜
*/
func stoneGame(piles []int) bool {

	//aMax bMin  aMin bMax        aMax+fMax(second)

	n := len(piles)

	s := 0

	for i := range piles {
		s += piles[i]
	}

	m1 := make([][]int, n)

	m2 := make([][]int, n)

	for i := range m1 {
		m1[i] = make([]int, n)
		m2[i] = make([]int, n)

		for j := range m2 {
			m1[i][j] = -1
			m2[i][j] = -1
		}
	}
	var first func(int, int) int

	var second func(int, int) int

	second = func(l int, r int) int {

		if l == r {
			return 0
		}

		if m2[l][r] != -1 {
			return m2[l][r]
		}

		cur := min(first(l+1, r), first(l, r-1))
		m2[l][r] = cur
		return cur
	}
	first = func(l, r int) int {
		if l == r {
			return piles[l]
		}
		if m1[l][r] != -1 {
			return m1[l][r]
		}
		cur := max(second(l+1, r)+piles[l], second(l, r-1)+piles[r])
		m1[l][r] = cur
		return cur
	}

	a := first(0, n-1)
	b := s - a

	return a > b
}

func stoneGame2(piles []int) bool {

	//aMax bMin  aMin bMax        aMax+fMax(second)

	return true
}
