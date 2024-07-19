package main

import (
	"fmt"
	"sort"
)

/**
给你一个二维整数数组 tiles ，其中 tiles[i] = [li, ri] ，表示所有在 li <= j <= ri 之间的每个瓷砖位置 j 都被涂成了白色。

同时给你一个整数 carpetLen ，表示可以放在 任何位置 的一块毯子的长度。

请你返回使用这块毯子，最多 可以盖住多少块瓷砖。


输入：tiles = [[1,5],[10,11],[12,18],[20,25],[30,32]], carpetLen = 10
输出：9
解释：将毯子从瓷砖 10 开始放置。
总共覆盖 9 块瓷砖，所以返回 9 。
注意可能有其他方案也可以覆盖 9 块瓷砖。
可以看出，瓷砖无法覆盖超过 9 块瓷砖。
*/

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {

	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	n := len(tiles)
	s := make([]int, n+1)
	//前缀和
	for i, v := range tiles {
		s[i+1] = s[i] + v[1] - v[0] + 1
	}
	ans := 0
	j := n - 1
	for i := n - 1; i >= 0; i-- {
		cur := tiles[i]
		//没超出边界可以继续往左扩
		for j > 0 && cur[1]-tiles[j][0]+1 < carpetLen {
			j--
		}
		//左边界
		mostCoverLeft := max(cur[1]-carpetLen+1, 0)
		cnt := s[i+1] - s[j+1] + max(tiles[j][1]-max(mostCoverLeft, tiles[j][0])+1, 0)
		ans = max(cnt, ans)
	}
	return ans

}

func main() {

	fmt.Println(maximumWhiteTiles([][]int{{8051, 8057}, {8074, 8089}, {7994, 7995}, {7969, 7987}, {8013, 8020}, {8123, 8139}, {7930, 7950}, {8096, 8104}, {7917, 7925}, {8027, 8035}, {8003, 8011}}, 9854))
}
