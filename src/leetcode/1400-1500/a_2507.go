package main

import "fmt"

/*
给你一个正整数 n 。

请你将 n 的值替换为 n 的 质因数 之和，重复这一过程。

注意，如果 n 能够被某个质因数多次整除，则在求和时，应当包含这个质因数同样次数。
返回 n 可以取到的最小值。

示例 1：

输入：n = 15
输出：5
解释：最开始，n = 15 。
15 = 3 * 5 ，所以 n 替换为 3 + 5 = 8 。
8 = 2 * 2 * 2 ，所以 n 替换为 2 + 2 + 2 = 6 。
6 = 2 * 3 ，所以 n 替换为 2 + 3 = 5 。
5 是 n 可以取到的最小值。
示例 2：

输入：n = 3
输出：3
解释：最开始，n = 3 。
3 是 n 可以取到的最小值。
*/
func smallestValue(n int) int {

	for {
		c := n
		s := 0
		for i := 2; i*i <= c; i++ {
			for c%i == 0 {
				s += i
				c /= i
			}
		}
		if c > 1 {
			s += c
		}

		if s == n {
			return n
		}
		n = s
	}
}

func main() {
	fmt.Println(smallestValue(4))
}
