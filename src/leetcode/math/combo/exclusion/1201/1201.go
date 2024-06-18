package main

import "fmt"

/*
*
丑数是可以被 a 或 b 或 c 整除的 正整数 。

给你四个整数：n 、a 、b 、c ，请你设计一个算法来找出第 n 个丑数。
*/
func nthUglyNumber(n int, a int, b int, c int) int {

	l, r := 0, n*min(a, b, c)
	lcmab := a * b / gcd(a, b)
	lcmac := a * c / gcd(a, c)
	lcmbc := b * c / gcd(b, c)
	lcmabc := a * lcmbc / (gcd(a, lcmbc))

	var check func(int) bool

	check = func(t int) bool {

		cnt := t/a + t/b + t/c - t/lcmab - t/lcmac - t/lcmbc + t/lcmabc
		return cnt >= n
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
	fmt.Println(nthUglyNumber(5, 2, 3, 3))
}
