package main

import (
	"fmt"
	"math"
)

func maximumAmount(coins [][]int) int {

	m := len(coins)
	n := len(coins[0])

	f := make([][][]int, m+1)

	for i := range f {

		f[i] = make([][]int, n+1)

		for j := range f[i] {
			f[i][j] = make([]int, 3)

			for k := range f[i][j] {

				f[i][j][k] = math.MinInt / 2
			}
		}

	}

	f[0][1][0] = 0
	f[1][0][0] = 0

	for i := 1; i <= m; i++ {

		for j := 1; j <= n; j++ {

			if coins[i-1][j-1] > 0 {

				for k := 0; k < 3; k++ {

					f[i][j][k] = max(f[i-1][j][k], f[i][j-1][k]) + coins[i-1][j-1]
				}

			} else {

				f[i][j][0] = max(f[i-1][j][0], f[i][j-1][0]) + coins[i-1][j-1]

				for k := 1; k < 3; k++ {
					f[i][j][k] = max(f[i-1][j][k-1], f[i][j-1][k-1], max(f[i-1][j][k], f[i][j-1][k])+coins[i-1][j-1])

				}

			}

		}
	}

	return max(f[m][n][2], f[m][n][1], f[m][n][0])

}

//[[-7,12,12,13],
//[-6,19,19,-6],
//[9,-2,-10,16],
//[-4,14,-10,-9]]

func main() {
	fmt.Println(maximumAmount([][]int{{-7, 12, 12, 13}, {-6, 19, 19, -6}, {9, -2, -10, 16}, {-4, 14, -10, -9}}))
}
