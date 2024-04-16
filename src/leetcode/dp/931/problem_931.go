//Given an n x n array of integers matrix, return the minimum sum of any
//falling path through matrix.
//
// A falling path starts at any element in the first row and chooses the
//element in the next row that is either directly below or diagonally left/right.
//Specifically, the next element from position (row, col) will be (row + 1, col - 1), (
//row + 1, col), or (row + 1, col + 1).
//
//
// Example 1:
//
//
//Input: matrix =
//[
//[2,1,3],
//[6,5,4],
//[7,8,9]
//]
//Output: 13
//Explanation: There are two falling paths with a minimum sum as shown.
//
//
// Example 2:
//
//
//Input: matrix = [[-19,57],[-40,-5]]
//Output: -59
//Explanation: The falling path with a minimum sum is shown.
//
//
//
// Constraints:
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 100
// -100 <= matrix[i][j] <= 100
//
//
// Related Topics Array Dynamic Programming Matrix 👍 5647 👎 136

package main

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minFallingPathSum(matrix [][]int) int {

	n := len(matrix)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
		dp[i][0] = matrix[0][i]
	}

	ans := math.MaxInt

	for i := 1; i < n; i++ {

		for j := 0; j < n; j++ {

			curMin := 0

			curMin = dp[j][(i-1)&1]

			if j-1 >= 0 {
				curMin = min(curMin, dp[j-1][(i-1)&1])
			}
			if j+1 < n {
				curMin = min(curMin, dp[j+1][(i-1)&1])
			}
			dp[j][i&1] = matrix[i][j] + curMin
			if i == n-1 {
				ans = min(dp[j][i&1], ans)
			}
		}
	}
	return ans
}

func min(a, b int) int {

	if a > b {
		return b
	}
	return a
}

func main() {
	println(minFallingPathSum([][]int{
		{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}

//leetcode submit region end(Prohibit modification and deletion)
