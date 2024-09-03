package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

/*
*
给你两个 正 整数 n 和 k 。

如果一个整数 x 满足以下条件，那么它被称为 k 回文 整数 。

x 是一个
回文整数 。
x 能被 k 整除。
如果一个整数的数位重新排列后能得到一个 k 回文整数 ，那么我们称这个整数为 好 整数。
比方说，k = 2 ，那么 2020 可以重新排列得到 2002 ，2002 是一个 k 回文串，所以 2020 是一个好整数。
而 1010 无法重新排列数位得到一个 k 回文整数。

请你返回 n 个数位的整数中，有多少个 好 整数。

4/2 =

思路：枚举+排列组合

注意 ，任何整数在重新排列数位之前或者之后 都不能 有前导 0 。比方说 1010 不能重排列得到 101 。
*/
func countGoodIntegers(n int, k int) int64 {

	m := (n - 1) / 2

	set := make(map[string]bool)
	s := int(math.Pow10(m))

	f := make([]int, n+1)

	f[0] = 1

	for i := 1; i <= n; i++ {
		f[i] = f[i-1] * i
	}
	ans := 0
	for i := s; i < s*10; i++ {

		x := i

		y := i

		if n%2 != 0 {
			y /= 10
		}

		for y > 0 {
			x = x*10 + y%10
			y /= 10
		}

		str := strconv.Itoa(x)
		if x%k != 0 {
			continue
		}

		bytes := []byte(str)
		slices.Sort(bytes)

		if set[string(bytes)] {
			continue
		}
		set[string(bytes)] = true
		cnt := make([]int, 10)
		for _, v := range str {
			cnt[v-'0'] += 1
		}
		cur := (n - cnt[0]) * f[n-1]
		for _, v := range cnt {
			cur /= f[v]
		}
		ans += cur
	}
	return int64(ans)
}

func main() {
	//fmt.Println(countGoodIntegers(3, 5))
	fmt.Println(countGoodIntegers(5, 6))
}
