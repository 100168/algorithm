package main

import (
	"fmt"
	"time"
)

/*
斐波那契数，通常用F(n) 表示，形成的序列称为 斐波那契数列 。该数列由0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1)= 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}
	if n == 2 {
		return 1
	}

	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		c := b
		b = a + b
		a = c
	}
	return b

}

func fib2(n int) int {

	p := [][]int{{1, 1}, {1, 0}}

	return pow(p, n-1)[0][0]

}

func pow(x [][]int, n int) [][]int {

	l := len(x)
	res := make([][]int, l)

	for i := range res {
		res[i] = make([]int, l)
		res[i][i] = 1
	}

	for ; n > 0; n >>= 1 {

		if n&1 != 0 {

			res = mul(res, x)
		}
		x = mul(x, x)
	}
	return res

}

func mul(a, b [][]int) [][]int {

	m, n := len(a), len(b[0])
	c := make([][]int, m)
	k := len(b)
	for i := range c {
		c[i] = make([]int, n)
	}

	for i := range c {
		for j := range c[i] {

			for p := 0; p < k; p++ {

				c[i][j] += a[i][p] * b[p][j]
			}
		}
	}
	return c
}

func main() {
	// 测量 fib 函数执行时间
	start1 := time.Now()
	result1 := fib(5000000000)
	duration1 := time.Since(start1)

	// 测量 fib2 函数执行时间
	start2 := time.Now()
	result2 := fib2(5000000000)
	duration2 := time.Since(start2)

	fmt.Printf("fib(500) = %d, 执行时间: %v\n", result1, duration1)
	fmt.Printf("fib2(500) = %d, 执行时间: %v\n", result2, duration2)

	if duration1 < duration2 {
		fmt.Println("fib 方法更快")
	} else {
		fmt.Println("fib2 方法更快")
	}
}
