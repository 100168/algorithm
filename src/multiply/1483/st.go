package main

import "math/bits"

func minValueInList(nums []int, op [][]int) []int {
	n := len(nums)
	m := bits.Len(uint(n))
	dp := make([][]int, n)
	for i := range nums {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}
	for j := 0; j < m-1; j++ {
		for x := 0; x < n; x++ {
			minIndex := min(x+1<<j, n-1)
			if nums[dp[x][j]] < nums[dp[minIndex][j]] {
				dp[x][j+1] = dp[x][j]
			} else {
				dp[x][j+1] = dp[minIndex][j]
			}
		}
	}

	ans := make([]int, len(op))
	for j, v := range op {
		index := v[0]
		curMin := nums[v[0]]
		for i := 0; i < n; i++ {
			if v[1]>>i&1 != 0 {
				index := dp[index][i]
				curMin = min(nums[index], curMin)
			}
		}
		ans[j] = curMin
	}
	return ans
}

func main() {
	println(minValueInList([]int{1, 5, 3, 6, 7, 2, 9, 10}, [][]int{{1, 8}}))
}
