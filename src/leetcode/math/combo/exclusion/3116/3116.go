package main

import (
	"fmt"
	"math/bits"
	"slices"
)

/*
*
给你一个整数数组 coins 表示不同面额的硬币，另给你一个整数 k 。

你有无限量的每种面额的硬币。但是，你 不能 组合使用不同面额的硬币。

返回使用这些硬币能制造的 第 kth 小 金额。
*/
func findKthSmallest(coins []int, k int) int64 {
	l, r := 1, slices.Min(coins)*k

	//预处理所有最小公倍数
	subsetLcm := make([]int, 1<<len(coins))
	subsetLcm[0] = 1
	for i, c := range coins {
		bit := 1 << i
		for mask, v := range subsetLcm[:bit] {
			subsetLcm[mask|bit] = lcm(v, c)
		}
	}

	var check func(int) bool
	check = func(n int) bool {
		cnt := 0
		for i := uint(1); i < 1<<len(coins); i++ {
			c := n / subsetLcm[i]
			if bits.OnesCount(i)%2 == 0 {
				c = -c
			}
			cnt += c

		}
		return cnt >= k
	}
	for l <= r {
		mid := (r-l)/2 + l

		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return int64(l)

}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	fmt.Println(findKthSmallest([]int{3, 6, 9}, 3))
}
