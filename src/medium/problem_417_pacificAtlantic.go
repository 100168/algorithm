package main

/*给定一个 m x n 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。

规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。

请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pacific-atlantic-water-flow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func pacificAtlantic(heights [][]int) [][]int {

	m := len(heights)
	n := len(heights[0])

	var ans [][]int

	p1 := make([][]bool, m)
	p2 := make([][]bool, m)
	for i := range p1 {
		p1[i] = make([]bool, n)
		p2[i] = make([]bool, n)

	}

	for i := 0; i < m; i++ {
		dfs(i, 0, p1, heights)
		dfs(i, n-1, p2, heights)
	}

	for i := 0; i < n; i++ {
		dfs(0, i, p1, heights)
		dfs(m-1, i, p2, heights)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p1[i][j] && p2[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func dfs(x int, y int, p [][]bool, heights [][]int) {

	m := len(heights)
	n := len(heights[0])

	//如果已经扫描过可达就不用扫描
	if p[x][y] {
		return
	}

	p[x][y] = true

	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, v := range directions {
		row := x + v[0]
		col := y + v[1]

		if row >= 0 && row < m && col >= 0 && col < n {
			if heights[x][y] <= heights[row][col] {
				dfs(row, col, p, heights)
			}
		}
	}

}
