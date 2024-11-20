package main

import "fmt"

/*
*
给你两个整数n 和k，和两个二维整数数组stayScore 和travelScore。

一位旅客正在一个有 n座城市的国家旅游，每座城市都 直接与其他所有城市相连。这位游客会旅游 恰好k天（下标从 0开始），
且旅客可以选择 任意城市作为起点。

每一天，这位旅客都有两个选择：

留在当前城市：如果旅客在第 i天停留在前一天所在的城市curr，旅客会获得stayScore[i][curr]点数。
前往另外一座城市：如果旅客从城市curr前往城市dest，旅客会获得travelScore[curr][dest]点数。
请你返回这位旅客可以获得的 最多点数。

示例 1：

输入：n = 2, k = 1, stayScore = [[2,3]], travelScore = [[0,2],[1,0]]

输出：3

解释：

旅客从城市 1 出发并停留在城市 1 可以得到最多点数。
*/
func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {

	f := make([][]int, n)

	ans := 0
	for i := range f {
		f[i] = make([]int, k)
	}
	for j := 0; j < k; j++ {
		for i := 0; i < n; i++ {
			f[i][j] = stayScore[j][i]
			if j > 0 {
				f[i][j] += f[i][j-1]
			}
			for from := 0; from < n; from++ {
				cur := travelScore[from][i]
				if j > 0 {
					cur += f[from][j-1]
				}
				f[i][j] = max(f[i][j], cur)
				ans = max(ans, f[i][j])
			}
		}
	}
	return ans

}

func main() {
	//fmt.Println(maxScore(2, 1, [][]int{{2, 3}}, [][]int{{0, 2}, {1, 0}}))
	//fmt.Println(maxScore(3, 2, [][]int{{3, 4, 2}, {2, 1, 2}}, [][]int{{0, 2, 1}, {2, 0, 4}, {3, 2, 0}}))
	//fmt.Println(maxScore(2, 1, [][]int{{1, 1}}, [][]int{{0, 1}, {6, 0}}))
	//fmt.Println(maxScore(2, 1, [][]int{{1, 1}}, [][]int{{0, 6}, {1, 0}}))
	//fmt.Println(maxScore(2, 1, [][]int{{1}, {1}}, [][]int{{0}}))
	fmt.Println(maxScore(2, 2, [][]int{{2, 5}, {4, 3}}, [][]int{{0, 3}, {4, 0}}))
}
