package main

import "fmt"

/*
*
请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。

让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。

由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。
*/
func numPrimeArrangements(n int) int {

	isPrime := make([]bool, n+1)

	mod := int(1e9 + 7)

	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	ans := 1
	cnt1 := 0
	cnt2 := 0
	for i := 1; i <= n; i++ {
		if isPrime[i] {
			cnt1++
			ans = ans * cnt1 % mod
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		} else {
			cnt2++
			ans = ans * cnt2 % mod
		}
	}
	return ans
}

func main() {
	fmt.Println(numPrimeArrangements(5))
}
