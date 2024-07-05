package main

/**
爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。

最初，黑板上有一个数字 n 。在每个玩家的回合，玩家需要执行以下操作：

选出任一 x，满足 0 < x < n 且 n % x == 0 。
用 n - x 替换黑板上的数字 n 。
如果玩家无法执行这些操作，就会输掉游戏。

只有在爱丽丝在游戏中取得胜利时才返回 true 。假设两个玩家都以最佳状态参与游戏。

1 a false
2 a true
3 a false
4 a true
5 a false
6 a true
7 a false
8 a true
9 a false

k+1

==>


*/

func divisorGame(n int) bool {

	return n%2 == 0
}
