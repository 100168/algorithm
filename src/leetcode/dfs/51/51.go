package main

import (
	"fmt"
	"math/bits"
	"slices"
	"strings"
)

/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
*/
func solveNQueens(n int) [][]string {

	ans := make([][]string, 0)
	size := 1<<n - 1

	var dfs func(int, int, int, []string)

	dfs = func(col int, left int, right int, path []string) {
		if col == size {
			ans = append(ans, slices.Clone(path))
		}

		pos := size &^ (col | left | right)
		for pos > 0 {
			p := pos & -pos
			pos -= p
			count := bits.TrailingZeros(uint(p))
			cur := strings.Repeat(".", n-count-1) + "Q" + strings.Repeat(".", count)
			path = append(path, cur)
			dfs(col|p, (left|p)<<1, (right|p)>>1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, 0, []string{})
	return ans
}

func main() {
	fmt.Println(solveNQueens(4))
}
