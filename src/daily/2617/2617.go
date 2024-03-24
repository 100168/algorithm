package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
)

func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colStacks := make([][]pair, n) // 每列的单调栈
	rowSt := []pair{}              // 行单调栈
	for i := m - 1; i >= 0; i-- {
		rowSt = rowSt[:0]
		for j := n - 1; j >= 0; j-- {
			colSt := colStacks[j]
			if i < m-1 || j < n-1 {
				mn = math.MaxInt
			}
			if g := grid[i][j]; g > 0 { // 可以向右/向下跳
				// 在单调栈上二分查找最优转移来源
				k := sort.Search(len(rowSt), func(k int) bool { return rowSt[k].i <= j+g })
				if k < len(rowSt) {
					mn = rowSt[k].x
				}
				k = sort.Search(len(colSt), func(k int) bool { return colSt[k].i <= i+g })
				if k < len(colSt) {
					mn = min(mn, colSt[k].x)
				}
			}
			if mn < math.MaxInt {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(rowSt) > 0 && mn <= rowSt[len(rowSt)-1].x {
					rowSt = rowSt[:len(rowSt)-1]
				}
				rowSt = append(rowSt, pair{mn, j})
				for len(colSt) > 0 && mn <= colSt[len(colSt)-1].x {
					colSt = colSt[:len(colSt)-1]
				}
				colStacks[j] = append(colSt, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt {
		return -1
	}
	return
}
func main() {
	//fmt.Println(minimumVisitedCells([][]int{{3, 4, 2, 1}, {4, 2, 3, 1}, {2, 1, 0, 0}, {2, 4, 0, 0}}))
	fmt.Println(minimumVisitedCells([][]int{{2, 1, 0}, {1, 0, 0}}))

	slices.Max()
}
