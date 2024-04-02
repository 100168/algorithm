package main

import (
	"math"
)

func minimumOperationsToWriteY(grid [][]int) int {

	n := len(grid)
	cntY := make([]int, 3)
	other := make([]int, 3)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (j == i && i <= n/2) || (i+j == n-1 && i <= n/2) || (j == n/2 && i >= n/2) {
				cntY[grid[i][j]]++
			} else {
				other[grid[i][j]]++
			}
		}
	}
	res := math.MaxInt

	sum1 := cntY[0] + cntY[1] + cntY[2]
	sumOthers := other[0] + other[1] + other[2]
	for i := range cntY {
		for j := range other {
			cur := sum1 - cntY[i]
			if j == i {
				continue
			}
			cur += sumOthers - other[j]
			res = min(res, cur)
		}
	}
	return res

}

func main() {

	print(minimumOperationsToWriteY([][]int{{1, 2, 2}, {1, 1, 0}, {0, 1, 0}}))
}
