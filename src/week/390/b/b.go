package main

import "math"

/*
暴力完事了
*/
func minOperations(k int) int {
	ans := math.MaxInt
	for i := 1; i <= k; i++ {
		cur := i - 1
		ops := (k - 1) / i
		ans = min(ans, cur+ops)
	}
	return ans
}
