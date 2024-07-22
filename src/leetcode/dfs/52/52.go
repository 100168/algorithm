package main

import "fmt"

/**
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

0001
0010
0000

0011
0110
0001
*/

func totalNQueens(n int) int {
	size := 1<<n - 1
	var dfs func(int, int, int) int
	dfs = func(col, left, right int) int {
		if col == size {
			return 1
		}
		cur := 0
		pos := size & (^(col | left | right))
		for pos > 0 {
			p := pos & -pos
			pos -= p
			cur += dfs(col|p, (left|p)>>1, (right|p)<<1)
		}
		return cur
	}
	return dfs(0, 0, 0)
}

func main() {
	fmt.Println(totalNQueens(14))
}
