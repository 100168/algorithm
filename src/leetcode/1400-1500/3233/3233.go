package main

import (
	"fmt"
	"math"
)

/**
给你两个 正整数 l 和 r。对于任何数字 x，x 的所有正因数（除了 x 本身）被称为 x 的 真因数。

如果一个数字恰好仅有两个 真因数，则称该数字为 特殊数字。例如：

数字 4 是 特殊数字，因为它的真因数为 1 和 2。
数字 6 不是 特殊数字，因为它的真因数为 1、2 和 3。
返回区间 [l, r] 内 不是 特殊数字 的数字数量。
*/

func nonSpecialCount(l int, r int) int {

	cnt := 0

	up := int(math.Sqrt(float64(r)))
	low := int(math.Sqrt(float64(l)))

	if low*low != l {
		low += 1
	}

	prime := make([]bool, up+1)

	for i := range prime {
		prime[i] = true
	}
	prime[1] = false
	for i := 2; i <= up; i++ {
		if prime[i] {
			for j := 2; j*i <= up; j++ {
				prime[j*i] = false
			}
		}
	}
	for i := low; i <= up; i++ {
		if prime[i] {
			cnt++
		}
	}
	return r - l + 1 - cnt
}

func main() {
	fmt.Println(nonSpecialCount(1, 2))
}
