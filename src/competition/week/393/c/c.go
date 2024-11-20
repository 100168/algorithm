package main

import (
	"math"
	"math/bits"
)

func findKthSmallest(coins []int, k int) int64 {
	l, r := int64(1), int64(math.MaxInt64)

	for l <= r {
		m := (r-l)/2 + l
		if check(coins, m, k) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return int64(l)
}

func gcd(a, b int64) int64 {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func check(nums []int, m int64, k int) bool {
	cnt := int64(0)
next:
	for i := uint(1); i < 1<<len(nums); i++ { // 枚举所有非空子集
		lcmRes := int64(1) // 计算子集 LCM
		for j := i; j > 0; j &= j - 1 {
			lcmRes = lcm(lcmRes, int64(nums[bits.TrailingZeros(j)]))
			if lcmRes > m { // 太大了
				goto next
			}
		}
		c := m / lcmRes
		if bits.OnesCount(i)%2 == 0 {
			c = -c
		}
		cnt += c
	}
	return cnt >= int64(k)
}

func main() {
	println(findKthSmallest([]int{5, 2}, 200000000))

}
