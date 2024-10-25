package main

import "fmt"

/*
*
给你三个整数 n ，x 和 y 。

一个活动总共有 n 位表演者。每一位表演者会 被安排 到 x 个节目之一，有可能有节目 没有 任何表演者。

所有节目都安排完毕后，评委会给每一个 有表演者的 节目打分，分数是一个 [1, y] 之间的整数。

请你返回 总 的活动方案数。

答案可能很大，请你将它对 109 + 7 取余 后返回。

注意 ，如果两个活动满足以下条件 之一 ，那么它们被视为 不同 的活动：

存在 一个表演者在不同的节目中表演。
存在 一个节目的分数不同。

示例 1：

输入：n = 1, x = 2, y = 3

输出：6

解释：

表演者可以在节目 1 或者节目 2 中表演。
评委可以给这唯一一个有表演者的节目打分 1 ，2 或者 3 。
示例 2：

输入：n = 5, x = 2, y = 2
//x*y
输出：32

解释：

每一位表演者被安排到节目 1 或者 2 。
所有的节目分数都为 1 。
示例 3：

输入：n = 3, x = 3, y = 4

输出：684

提示：

1 <= n, x, y <= 1000
*/

const mod = 1_000_000_007
const mx = 1001

var s [mx][mx]int

func init() {
	s[0][0] = 1
	for i := 1; i < mx; i++ {
		for j := 1; j <= i; j++ {
			s[i][j] = (s[i-1][j-1] + j*s[i-1][j]) % mod
		}
	}
}

func numberOfWays(n, x, y int) (ans int) {
	perm, powY := 1, 1
	for i := 1; i <= min(n, x); i++ {
		perm = perm * (x + 1 - i) % mod
		powY = powY * y % mod
		ans = (ans + perm*s[n][i]%mod*powY) % mod
	}
	return
}

func main() {

	//3*3*3
	fmt.Println(numberOfWays(2, 2, 2))
}
