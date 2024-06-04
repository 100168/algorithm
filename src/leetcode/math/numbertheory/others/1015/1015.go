package main

import "fmt"

/*
*
给定正整数 k ，你需要找出可以被 k 整除的、仅包含数字 1 的最 小 正整数 n 的长度。

返回 n 的长度。如果不存在这样的 n ，就返回-1。

注意： n 可能不符合 64 位带符号整数。

思路：同余
*/
func smallestRepunitDivByK(k int) int {

	if k&1 == 0 {
		return -1
	}

	mark := make(map[int]bool)

	cnt := 1
	for rest := 1; rest%k != 0; rest = (rest*10 + 1) % k {
		if mark[rest] {
			return -1
		}
		mark[rest] = true
		cnt++
	}
	return cnt
}

// 计算欧拉函数（n 以内的与 n 互质的数的个数）
func phi(n int) int {
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = res / i * (i - 1)
			for n /= i; n%i == 0; n /= i {
			}
		}
	}
	if n > 1 {
		res = res / n * (n - 1)
	}
	return res
}

// 快速幂，返回 pow(x, n) % mod
func pow(x, n, mod int) int {
	x %= mod
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func smallestRepunitDivByK3(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	m := phi(k * 9)
	// 从小到大枚举不超过 sqrt(m) 的因子
	i := 1
	for ; i*i <= m; i++ {
		if m%i == 0 && pow(10, i, k*9) == 1 {
			return i
		}
	}
	// 从小到大枚举不低于 sqrt(m) 的因子
	for i--; ; i-- {
		if m%i == 0 && pow(10, m/i, k*9) == 1 {
			return m / i
		}
	}
}

func main() {
	fmt.Println(smallestRepunitDivByK(1))
}
