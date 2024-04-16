package main

func pathsWithMaxScore(board []string) []int {

	m := len(board)
	n := len(board[0])

	type pair struct{ val, cnt int }

	dp := make([][]pair, m)

	for i := range dp {
		dp[i] = make([]pair, n)
	}
	dp[m-1][n-1] = pair{0, 1}
	for i := m - 2; i >= 0; i-- {
		if board[i][n-1] != 'X' && dp[i+1][n-1].val != -1 {
			dp[i][n-1] = pair{dp[i+1][n-1].val + int(board[i][n-1]-'0'), 1}
		} else {
			dp[i][n-1] = pair{-1, 1}
		}
	}
	for j := n - 2; j >= 0; j-- {
		if board[m-1][j] != 'X' && dp[m-1][j+1].val != -1 {
			dp[m-1][j] = pair{dp[m-1][j+1].val + int(board[m-1][j]-'0'), 1}
		} else {
			dp[m-1][j] = pair{-1, 1}
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			cur := board[i][j]
			if cur == 'X' {
				dp[i][j] = pair{-1, 1}
				continue
			}
			dp[i][j].val = max(dp[i+1][j].val, dp[i+1][j+1].val, dp[i][j+1].val)

			if dp[i][j].val == -1 {
				continue
			}
			if dp[i+1][j].val == dp[i][j].val {
				dp[i][j].cnt += dp[i+1][j].cnt
			}
			if dp[i+1][j+1].val == dp[i][j].val {
				dp[i][j].cnt += dp[i+1][j+1].cnt
			}
			if dp[i][j+1].val == dp[i][j].val {
				dp[i][j].cnt += dp[i][j+1].cnt
			}
			if i == 0 && j == 0 {
				break
			}
			dp[i][j].val += int(cur - '0')
		}
	}
	if dp[0][0].val == -1 {
		return []int{0, 0}
	}
	return []int{dp[0][0].val, dp[0][0].cnt}
}

func main() {
	println(pathsWithMaxScore([]string{"EX", "XS"}))
}
