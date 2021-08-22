package main

/*给定一个包含了一些 0 和 1 的非空二维数组grid 。

一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func maxAreaOfIsland(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	maxLength := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				temp := infect(grid, i, j, n-1, m-1)
				if temp > maxLength {
					maxLength = temp
				}
			}
		}
	}
	return maxLength
}

func infect(grid [][]int, sr int, sc int, n int, m int) int {
	res := 0
	if grid[sr][sc] == 1 {
		grid[sr][sc] = 2
		res = 1
	} else {
		return 0
	}
	if sr+1 <= n {
		res = res + infect(grid, sr+1, sc, n, m)
	}
	if sr-1 >= 0 {
		res = res + infect(grid, sr-1, sc, n, m)
	}
	if sc+1 <= m {
		res = res + infect(grid, sr, sc+1, n, m)
	}
	if sc-1 >= 0 {
		res = res + infect(grid, sr, sc-1, n, m)
	}
	return res
}

func main() {

}
