package main

import "fmt"

/**
给你一个二维整数数组 grid ，大小为 m x n，其中每个单元格都含一个正整数。

转角路径 定义为：包含至多一个弯的一组相邻单元。具体而言，路径应该完全 向水平方向 或者 向竖直方向 移动过弯（如果存在弯），而不能访问之前访问过的单元格。在过弯之后，路径应当完全朝 另一个 方向行进：如果之前是向水平方向，那么就应该变为向竖直方向；反之亦然。当然，同样不能访问之前已经访问过的单元格。

一条路径的 乘积 定义为：路径上所有值的乘积。

请你从 grid 中找出一条乘积中尾随零数目最多的转角路径，并返回该路径中尾随零的数目。

注意：

水平 移动是指向左或右移动。
竖直 移动是指向上或下移动。

2,5,10

4 * 25

8 * 125

*/

func maxTrailingZeros(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	rMulLeft := make([][]pair, m)
	colMulUp := make([][]pair, m)
	rMulRight := make([][]pair, m)
	colMulDown := make([][]pair, m)

	for i := 0; i < m; i++ {
		rMulLeft[i] = make([]pair, n)
		colMulUp[i] = make([]pair, n)
		rMulRight[i] = make([]pair, n)
		colMulDown[i] = make([]pair, n)
	}

	for i := range grid {
		for j := range grid[i] {
			if i > 0 {
				pre := colMulUp[i-1][j]
				colMulUp[i][j] = getPair(pre, grid[i][j])
			} else {
				pre := pair{0, 0}
				colMulUp[i][j] = getPair(pre, grid[i][j])
			}
			if j > 0 {
				pre := rMulLeft[i][j-1]
				rMulLeft[i][j] = getPair(pre, grid[i][j])
			} else {
				pre := pair{0, 0}
				rMulLeft[i][j] = getPair(pre, grid[i][j])
			}
		}
		for j := n - 1; j >= 0; j-- {
			if j < n-1 {
				pre := rMulRight[i][j+1]
				rMulRight[i][j] = getPair(pre, grid[i][j])
			} else {
				pre := pair{0, 0}
				rMulRight[i][j] = getPair(pre, grid[i][j])
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := range grid[i] {
			if i < m-1 {
				pre := colMulDown[i+1][j]
				colMulDown[i][j] = getPair(pre, grid[i][j])
			} else {
				pre := pair{0, 0}
				colMulDown[i][j] = getPair(pre, grid[i][j])
			}
		}
	}
	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			curRL := rMulLeft[i][j]
			curRR := rMulRight[i][j]
			curCUp := pair{0, 0}
			curCDown := pair{0, 0}
			if i+1 < m {
				curCDown = colMulDown[i+1][j]
			}
			if i > 0 {
				curCUp = colMulUp[i-1][j]
			}
			cur := max(getRes(curRL, curCDown), getRes(curRL, curCUp), getRes(curRR, curCDown), getRes(curRR, curCUp))
			ans = max(ans, cur)
		}
	}
	return ans
}

func getRes(pre, cur pair) int {

	return min(pre.cntTwo+cur.cntTwo, pre.cntFive+cur.cntFive)
}

type pair struct {
	cntTwo  int
	cntFive int
}

func getPair(pre pair, cur int) pair {

	cntTwo := pre.cntTwo
	cntFive := pre.cntFive
	for cur%2 == 0 {
		cur /= 2
		cntTwo++
	}
	for cur%5 == 0 {
		cntFive++
		cur /= 5
	}
	return pair{cntTwo, cntFive}
}

func main() {
	fmt.Println(maxTrailingZeros([][]int{{23, 17, 15, 3, 20}, {8, 1, 20, 27, 11}, {9, 4, 6, 2, 21}, {40, 9, 1, 10, 6}, {22, 7, 4, 5, 3}}))
	fmt.Println(maxTrailingZeros([][]int{{1, 5, 2, 4, 25}}))
}
