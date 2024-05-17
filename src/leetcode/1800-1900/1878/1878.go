package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*
菱形和 指的是 grid 中一个正菱形 边界 上的元素之和。本题中的菱形必须为正方形旋转45度，
且四个角都在一个格子当中。下图是四个可行的菱形，每个菱形和应该包含的格子都用了相应颜色标注在图中。
注意，菱形可以是一个面积为 0 的区域，如上图中右下角的紫色菱形所示。
请你按照 降序 返回 grid 中三个最大的 互不相同的菱形和 。如果不同的和少于三个，则将它们全部返回。
todo mark
*/
func getBiggestThree(grid [][]int) []int {

	m, n := len(grid), len(grid[0])

	var ans []int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = append(ans, grid[i][j])
			for k := 1; k < min(m, n); k++ {
				if j+k >= n || i+k >= m || j-k < 0 || i+2*k >= m {
					continue
				}
				cur := 0
				for l := 0; l <= k; l++ {
					cur += grid[i+l][j-l]
					cur += grid[i+l][j+l]
					cur += grid[i+2*k-l][j-l]
					cur += grid[i+2*k-l][j+l]
				}
				cur -= grid[i][j] + grid[i+2*k][j] + grid[i+k][j-k] + grid[i+k][j+k]
				ans = append(ans, cur)
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	ans = slices.Compact(ans)
	if len(ans) <= 3 {
		return ans
	}
	ans = ans[:3]
	return ans
}

func getBiggestThree2(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	ds := make([][]int, n+1) // 主对角线方向前缀和
	as := make([][]int, n+1) // 反对角线方向前缀和
	for i := range ds {
		ds[i] = make([]int, m+1)
		as[i] = make([]int, m+1)
	}
	for i, r := range grid {
		for j, v := range r {
			ds[i+1][j+1] = ds[i][j] + v // ↘
			as[i+1][j] = as[i][j+1] + v // ↙
		}
	}
	// 从 (x,y) 开始，向 ↘，连续的 k 个数的和
	queryDiagonal := func(x, y, k int) int { return ds[x+k][y+k] - ds[x][y] }
	// 从 (x,y) 开始，向 ↙，连续的 k 个数的和
	queryAntiDiagonal := func(x, y, k int) int { return as[x+k][y+1-k] - as[x][y+1] }

	var x, y, z int // 最大，严格次大，严格第三大
	update := func(v int) {
		if v > x {
			x, y, z = v, x, y
		} else if v < x && v > y {
			y, z = v, y
		} else if v < y && v > z {
			z = v
		}
	}

	for i, r := range grid {
		for j, v := range r {
			update(v)
			for k := 1; k <= i && i+k < n && k <= j && j+k < m; k++ {
				a := queryDiagonal(i-k, j, k)           // 菱形右上
				b := queryAntiDiagonal(i-k+1, j-1, k-1) // 菱形左上
				c := queryDiagonal(i, j-k, k)           // 菱形左下
				d := queryAntiDiagonal(i, j+k, k+1)     // 菱形右下
				update(a + b + c + d)
			}
		}
	}

	ans := []int{x, y, z}
	for ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	return ans
}
func main() {
	fmt.Println(getBiggestThree([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))
	fmt.Println(getBiggestThree2([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}))
}
