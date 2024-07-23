package main

import "math"

/*
*
你有一大块巧克力，它由一些甜度不完全相同的小块组成。我们用数组 sweetness 来表示每一小块的甜度。

你打算和 K 名朋友一起分享这块巧克力，所以你需要将切割 K 次才能得到 K+1 块，每一块都由一些 连续 的小块组成。

为了表现出你的慷慨，你将会吃掉 总甜度最小 的一块，并将其余几块分给你的朋友们。

请找出一个最佳的切割策略，使得你所分得的巧克力 总甜度最大，并返回这个 最大总甜度。
*/
func maximizeSweetness(sweetness []int, k int) int {
	l, r := 0, math.MaxInt/2

	check := func(t int) bool {
		s := 0
		cnt := 0
		for _, v := range sweetness {
			s += v
			if s >= t {
				cnt++
				s = 0
			}
		}
		if s >= t {
			cnt++
		}
		return cnt > k
	}

	for l <= r {

		m := (r-l)/2 + l

		if check(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
