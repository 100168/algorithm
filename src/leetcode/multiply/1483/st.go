package main

import "math/bits"

type St struct {
	dp [][]int
}

func constructor(nums []int) St {
	st := new(St)
	m := len(nums)
	n := bits.Len(uint(m))
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = nums[i]
	}
	st.dp = dp
	for j := 0; j < n-1; j++ {
		for i := 0; i+1<<j < m; i++ {
			dp[i][j+1] = st.strategy(i, j)
		}
	}
	return *st
}

func (st *St) strategy(i int, j int) int {
	return min(st.dp[i][j], st.dp[i+1<<j][j])
}
func (st *St) query(l, r int) int {
	k := 0
	for ; 1<<(k+1) <= r-l+1; k++ {
	}
	return st.queryLR(l, r, k)
}
func (st *St) queryLR(l, r, k int) int {
	return min(st.dp[l][k], st.dp[r-1<<k+1][k])
}

func main() {
	//dp[5][1] = dp[5][0] dp[5+1][0]
	st := constructor([]int{2, 2, 5, 4, 4, 1})
	//                      0  1  2  3  4  5
	println(st.query(0, 5))

	/**
	2 2 2
	2 2 2
	5 4 1
	4 4 0
	4 1 0
	1 0 0

	dp[5][1] = 0

	*/
}
