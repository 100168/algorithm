package main

import (
	"fmt"
	"math"
)

/*
*
一个正整数如果能被 a 或 b 整除，那么它是神奇的。
给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 109 + 7 取模 后的值。
*/
func nthMagicalNumber(n int, a int, b int) int {

	mod := int(1e9 + 7)
	var check func(int) bool
	g := gcd(a, b)
	lcm := a * b / g
	check = func(t int) bool {
		cnt := t/a + t/b - t/lcm
		return cnt >= n
	}

	l, r := 0, math.MaxInt

	for l <= r {
		m := (r-l)/2 + l
		if check(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l % mod

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(nthMagicalNumber(5, 7, 5))
}
