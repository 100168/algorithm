package main

import (
	"fmt"
	"math"
)

/*
*

几块石子 排成一行 ，每块石子都有一个关联值，关联值为整数，由数组 stoneValue 给出。
游戏中的每一轮：Alice 会将这行石子分成两个 非空行（即，左侧行和右侧行）；
Bob 负责计算每一行的值，即此行中所有石子的值的总和。
Bob 会丢弃值最大的行，Alice 的得分为剩下那行的值（每轮累加）。
如果两行的值相等，Bob 让 Alice 决定丢弃哪一行。下一轮从剩下的那一行开始。

只 剩下一块石子 时，游戏结束。Alice 的分数最初为 0 。

返回 Alice 能够获得的最大分数

输入：stoneValue = [6,2,3,4,5,5]
输出：18
解释：在第一轮中，Alice 将行划分为 [6，2，3]，[4，5，5] 。左行的值是 11 ，右行的值是 14 。Bob 丢弃了右行，Alice 的分数现在是 11 。
在第二轮中，Alice 将行分成 [6]，[2，3] 。这一次 Bob 扔掉了左行，Alice 的分数变成了 16（11 + 5）。
最后一轮 Alice 只能将行分成 [2]，[3] 。Bob 扔掉右行，Alice 的分数现在是 18（16 + 2）。游戏结束，因为这行只剩下一块石头了。

思路；

按题意直接区间dp
*/
func stoneGameV(stoneValue []int) int {

	n := len(stoneValue)

	s := make([]int, n+1)
	for i, v := range stoneValue {
		s[i+1] = s[i] + v
	}
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(l, r int) int {
		if l == r {
			return 0
		}

		if f[l][r] != -1 {
			return f[l][r]
		}

		cur := math.MinInt

		for m := l; m < r; m++ {
			lS := s[m+1] - s[l]
			rS := s[r+1] - s[m+1]
			if lS == rS {
				cur = max(cur, dfs(l, m)+lS, dfs(m+1, r)+rS)
			} else if lS < rS {
				cur = max(cur, dfs(l, m)+lS)
			} else {
				cur = max(cur, dfs(m+1, r)+rS)
			}
		}
		f[l][r] = cur
		return cur

	}

	return dfs(0, n-1)

}

func main() {

	//
	fmt.Println(stoneGameV([]int{1, 2, 3, 1}))
}
