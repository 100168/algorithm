package main

import (
	"fmt"
	"math"
)

/*
给你两个 正整数 l 和 r。对于任何数字 x，x 的所有正因数（除了 x 本身）被称为 x 的 真因数。

如果一个数字恰好仅有两个 真因数，则称该数字为 特殊数字。例如：

数字 4 是 特殊数字，因为它的真因数为 1 和 2。
数字 6 不是 特殊数字，因为它的真因数为 1、2 和 3。
返回区间 [l, r] 内 不是 特殊数字 的数字数量。

*/

func nonSpecialCount(l int, r int) int {

	s := int(math.Sqrt(float64(l)))

	if s*s != l {
		s += 1
	}

	end := int(math.Sqrt(float64(r)))

	isPrime := make([]bool, end+1)

	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	for i := 2; i <= end; i++ {

		if !isPrime[i] {
			continue
		}

		for j := 2; j*i <= end; j++ {
			isPrime[j*i] = false
		}
	}

	ans := 0
	for i := s; i <= end; i++ {
		if isPrime[i] {
			ans++
		}

	}
	return r - l + 1 - ans

}

func main() {
	//fmt.Println(nonSpecialCount(5, 7))
	fmt.Println(nonSpecialCount(4, 16))
}
