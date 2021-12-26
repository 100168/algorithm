package main

import "fmt"

/*假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。*/

func climbStairs(n int) int {

	a := 1
	b := 1
	for i := 2; i <= n; i++ {
		c := b
		b = a + b
		a = c
	}
	return b
}
func main() {

	stairs := climbStairs(11)
	fmt.Println(stairs)
}
