package main

import (
	"fmt"
	"math"
)

/*
*

给你一个二维 二进制 数组 grid。你需要找到 3 个 不重叠、面积 非零 、
边在水平方向和竖直方向上的矩形，并且满足 grid 中所有的 1 都在这些矩形的内部。

返回这些矩形面积之和的 最小 可能值。

注意，这些矩形可以相接。
*/
func minimumSum(grid [][]int) int {

	ans := min(getMinSum(grid), getMinSum(rotation(grid)))
	return ans
}

func getMinSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	var getMinArea func(r1, c1, r2, c2 int) int
	getMinArea = func(r1, c1, r2, c2 int) int {

		up := r2
		down := r1
		left := c2
		right := c1
		s := 0
		for i := r1; i <= r2; i++ {
			for j := c1; j <= c2; j++ {
				if grid[i][j] == 1 {
					up = min(up, i)
					down = max(down, i)
					left = min(left, j)
					right = max(right, j)
					s++
				}
			}
		}

		if s == 0 {
			return 1
		}
		return (right - left + 1) * (down - up + 1)
	}
	ans := math.MaxInt
	for i := 0; i < m-1; i++ {
		area1 := getMinArea(0, 0, i, n-1)
		for j := 0; j < n-1; j++ {
			area2 := getMinArea(i+1, 0, m-1, j)
			area3 := getMinArea(i+1, j+1, m-1, n-1)
			ans = min(ans, area1+area2+area3)
		}
	}

	for i := 0; i < m-2; i++ {
		area1 := getMinArea(0, 0, i, n-1)
		for j := i + 1; j < m-1; j++ {
			area2 := getMinArea(i+1, 0, j, n-1)
			area3 := getMinArea(j+1, 0, m-1, n-1)
			ans = min(ans, area1+area2+area3)
		}
	}

	for i := 1; i < m; i++ {
		area1 := getMinArea(i, 0, m-1, n-1)
		for j := 0; j < n-1; j++ {
			area2 := getMinArea(0, 0, i-1, j)
			area3 := getMinArea(0, j+1, i-1, n-1)
			ans = min(ans, area1+area2+area3)
		}

	}
	return ans

}

func rotation(grid [][]int) [][]int {

	m := len(grid)
	n := len(grid[0])
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			g[n-j-1][i] = grid[i][j]
		}
	}
	return g
}

func main() {
	grid := [][]int{
		{1, 1},
		{0, 1},
	}
	fmt.Println("Minimum possible sum of areas:", minimumSum(grid)) // Output: 8
}
