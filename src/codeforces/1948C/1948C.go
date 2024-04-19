package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1948C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var t, n int

	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		fmt.Fscan(in, &n)
		grid := make([]string, 2)
		for i := range grid {
			fmt.Fscan(in, &grid[i])
		}

		dp := make([][][]bool, 2)

		for i := range dp {
			dp[i] = make([][]bool, n)
			for j := range dp[i] {
				dp[i][j] = make([]bool, 2)
			}
		}

		//表示当前不能用
		dp[0][0][1] = true
		//表示单前
		dp[1][0][0] = true
		//>>><
		//>><<
		for j := 1; j < n; j++ {

			//当前可以换方向
			dp[0][j][1] = dp[0][j-1][0] && grid[0][j-1] == '>'
			dp[1][j][1] = dp[1][j-1][0] && grid[1][j-1] == '>'

			//当前不能换方向
			dp[0][j][0] = dp[0][j-1][1] || dp[1][j-1][0] && grid[1][j-1] == '>'
			dp[1][j][0] = dp[1][j-1][1] || dp[0][j-1][0] && grid[0][j-1] == '>'

		}

		if dp[1][n-1][0] || dp[1][n-1][1] {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func main() {
	CF1948C(os.Stdin, os.Stdout)
}
