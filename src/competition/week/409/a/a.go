package main

/*
*
给你一个 n x n 的二维数组 grid，它包含范围 [0, n2 - 1] 内的不重复元素。

实现 neighborSum 类：

neighborSum(int [][]grid) 初始化对象。
int adjacentSum(int value) 返回在 grid 中与 value 相邻的元素之和，相邻指的是与 value 在上、左、右或下的元素。
int diagonalSum(int value) 返回在 grid 中与 value 对角线相邻的元素之和，对角线相邻指的是与 value 在左上、右上、左下或右下的元素。
*/
type neighborSum struct {
	indexMap map[int]int
	grid     [][]int
}

func Constructor(grid [][]int) neighborSum {

	ng := new(neighborSum)
	ng.grid = grid
	ng.indexMap = make(map[int]int)
	for i, r := range grid {
		for j, v := range r {
			ng.indexMap[v] = i*len(grid[0]) + j
		}
	}
	return *ng
}

func (t *neighborSum) AdjacentSum(value int) int {

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	index := t.indexMap[value]
	r := index / len(t.grid)
	c := index % len(t.grid[0])

	ans := 0
	for _, d := range dir {

		x, y := r+d[0], c+d[1]

		if x < len(t.grid) && x >= 0 && y < len(t.grid[0]) && y >= 0 {
			ans += t.grid[x][y]
		}
	}
	return ans
}

func (t *neighborSum) DiagonalSum(value int) int {
	dir := [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	index := t.indexMap[value]
	r := index / len(t.grid)
	c := index % len(t.grid[0])
	ans := 0
	for _, d := range dir {
		x, y := r+d[0], c+d[1]

		if x < len(t.grid) && x >= 0 && y < len(t.grid[0]) && y >= 0 {
			ans += t.grid[x][y]
		}
	}
	return ans
}

func main() {

}
