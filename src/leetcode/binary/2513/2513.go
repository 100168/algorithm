package main

import "math"

/*
*
给你两个数组 arr1 和 arr2 ，它们一开始都是空的。你需要往它们中添加正整数，使它们满足以下条件：

arr1 包含 uniqueCnt1 个 互不相同 的正整数，每个整数都 不能 被 divisor1 整除 。
arr2 包含 uniqueCnt2 个 互不相同 的正整数，每个整数都 不能 被 divisor2 整除 。
arr1 和 arr2 中的元素 互不相同 。
给你 divisor1 ，divisor2 ，uniqueCnt1 和 uniqueCnt2 ，请你返回两个数组中 最大元素 的 最小值 。

思路：二分+容斥原理

1.
*/
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	l, r := 0, math.MaxInt

	g := gcd(divisor1, divisor2)
	lcm := divisor1 * divisor2 / g
	var check func(int) bool
	check = func(n int) bool {
		cnt := n - max(n/divisor1-n/lcm-uniqueCnt2, 0) - max(n/divisor2-n/lcm-uniqueCnt1, 0) - n/lcm
		return cnt >= uniqueCnt1+uniqueCnt2
	}
	for l <= r {
		m := (r-l)/2 + l
		if check(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	println(minimizeSet(2, 7, 1, 3))
	println(minimizeSet(2, 4, 8, 2))
}
