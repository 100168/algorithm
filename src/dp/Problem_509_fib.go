package main

import "fmt"

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

func main() {
	fmt.Println("hello")
}
