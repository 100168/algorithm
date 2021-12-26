package main

import "fmt"

/*一个 2D 网格中的 顶峰元素 是指那些 严格大于 其相邻格子(上、下、左、右)的元素。

给你一个 从 0 开始编号 的 m x n 矩阵 mat ，其中任意两个相邻格子的值都 不相同 。找出 任意一个 顶峰元素 mat[i][j] 并 返回其位置 [i,j] 。

你可以假设整个矩阵周边环绕着一圈值为 -1 的格子。

要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-a-peak-element-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
func findPeakGrid(mat [][]int) []int {

	n := len(mat[0])
	l := 0
	r := len(mat) - 1

	//按行为维度进行二分，找出每行最大值，
	for l < r {
		middle := l + (r-l)>>1
		middleMax := getMax(middle, n-1, mat)
		leftMax := getMax(l, n-1, mat)
		rightMax := getMax(r, n-1, mat)
		if mat[middle][middleMax] > mat[l][leftMax] && mat[middle][middleMax] > mat[r][rightMax] {
			return []int{middle, middleMax}
		} else if mat[l][leftMax] > mat[r][rightMax] {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}
	max := getMax(l, n-1, mat)

	return []int{l, max}
}

func getMax(x int, n int, mat [][]int) int {
	maxCol := 0

	for i := 1; i <= n; i++ {
		if mat[x][i] > mat[x][maxCol] {
			maxCol = i
		}
	}
	return maxCol
}
func main() {
	mat := [][]int{{1, 2}, {2, 3}}

	grid := findPeakGrid(mat)
	fmt.Println(grid)
}
