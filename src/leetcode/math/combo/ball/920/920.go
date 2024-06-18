package main

import "fmt"

/*
*
你的音乐播放器里有 n 首不同的歌，在旅途中，你计划听 goal 首歌（不一定不同，即，允许歌曲重复）。你将会按如下规则创建播放列表：

每首歌 至少播放一次 。
一首歌只有在其他 k 首歌播放完之后才能再次播放。
给你 n、goal 和 k ，返回可以满足要求的播放列表的数量。由于答案可能非常大，请返回对 109 + 7 取余 的结果。

思路：dp
1.状态定义：dp[i][j]  当前听了i首歌，并且有首不同，
2.状态转移：
 1. 当前选择的是前面都没有选过的：dp[i][j] = dp[i-1][j-1]*(n-(j-1))
 2. 当前选择的是前面选过的：dp[i][j] = dp[i-1][j]*max(j-k,0)
*/
func numMusicPlaylists(n int, goal int, k int) int {

	mod := int(1e9 + 7)
	dp := make([][]int, goal+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= goal; i++ {
		for j := 1; j <= n && j <= i; j++ {
			dp[i][j] = (dp[i-1][j-1]*(n-j+1) + dp[i-1][j]*max(j-k, 0)) % mod
		}
	}
	return dp[goal][n]
}

func main() {
	fmt.Println(numMusicPlaylists(3, 3, 1))
}
