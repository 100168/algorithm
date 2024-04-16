package main

//给你一个大小为 m x n 的矩阵 grid 。最初，你位于左上角 (0, 0) ，每一步，你可以在矩阵中 向右 或 向下 移动。
//
// 在从左上角 (0, 0) 开始到右下角 (m - 1, n - 1) 结束的所有路径中，找出具有 最大非负积 的路径。路径的积是沿路径访问的单元格中所有整
//数的乘积。
//
// 返回 最大非负积 对 10⁹ + 7 取余 的结果。如果最大积为 负数 ，则返回 -1 。
//
// 注意，取余是在得到最大积之后执行的。
//
//
//
// 示例 1：
//
//
//输入：grid = [[-1,-2,-3],[-2,-3,-3],[-3,-3,-2]]
//输出：-1
//解释：从 (0, 0) 到 (2, 2) 的路径中无法得到非负积，所以返回 -1 。
//
// 示例 2：
//
//
//输入：grid = [[1,-2,1],[1,-2,1],[3,-4,1]]
//输出：8
//解释：最大非负积对应的路径如图所示 (1 * 1 * -2 * -4 * 1 = 8)
//
//
// 示例 3：
//
//
//输入：grid = [[1,3],[0,-4]]
//输出：0
//解释：最大非负积对应的路径如图所示 (1 * 0 * -4 = 0)
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 15
// -4 <= grid[i][j] <= 4
//
//
// Related Topics 数组 动态规划 矩阵 👍 50 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maxProductPath(grid [][]int) int {

	mod := int64(1_000_000_007)
	m := len(grid)
	n := len(grid[0])
	dpMin := make([][]int64, m)
	dpMax := make([][]int64, m)
	for i := range dpMax {
		dpMax[i] = make([]int64, n)
		dpMin[i] = make([]int64, n)
	}

	dpMin[0][0] = int64(grid[0][0])
	dpMax[0][0] = int64(grid[0][0])
	for i := 1; i < m; i++ {
		dpMin[i][0] = (dpMin[i-1][0] * int64(grid[i][0])) % mod
		dpMax[i][0] = (dpMax[i-1][0] * int64(grid[i][0])) % mod

	}
	for j := 1; j < n; j++ {
		dpMin[0][j] = (dpMin[0][j-1] * int64(grid[0][j])) % mod
		dpMax[0][j] = (dpMax[0][j-1] * int64(grid[0][j])) % mod
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dpMin[i][j] = min(dpMax[i-1][j]*int64(grid[i][j]), dpMax[i][j-1]*int64(grid[i][j]), dpMin[i-1][j]*int64(grid[i][j]), dpMin[i][j-1]*int64(grid[i][j]))
			dpMax[i][j] = max(dpMin[i-1][j]*int64(grid[i][j]), dpMin[i][j-1]*int64(grid[i][j]), dpMax[i-1][j]*int64(grid[i][j]), dpMax[i][j-1]*int64(grid[i][j]))
		}
	}
	if dpMax[m-1][n-1] < 0 {
		return -1
	}
	return int(dpMax[m-1][n-1] % mod)
}

func main() {

	println(maxProductPath([][]int{
		{-1, -4, 2},
		{4, 3, -1},
		{2, -4, 4},
		{1, -1, -4}}))
	println((-4) * 4 * (-4) * 3 * (-4) * (-1))
	println(4 * (-4) * 3 * (-4) * (-1))
}

//leetcode submit region end(Prohibit modification and deletion)
