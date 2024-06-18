package main

import "fmt"

/*
*
给你 3 个正整数 zero ，one 和 limit 。

一个二进制数组arr 如果满足以下条件，那么我们称它是 稳定的 ：

0 在 arr 中出现次数 恰好 为 zero 。
1 在 arr 中出现次数 恰好 为 one 。
arr 中每个长度超过 limit 的
子数组都 同时 包含 0 和 1 。

请你返回 稳定 二进制数组的 总 数目。

由于答案可能很大，将它对 109 + 7 取余 后返回。
*/
const mod = 1000_000_007
const mx = 1001

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
}

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func numberOfStableArrays(zero, one, limit int) (ans int) {
	if zero > one {
		zero, one = one, zero // 保证空间复杂度为 O(min(zero, one))
	}
	f0 := make([]int, zero+3)
	for i := (zero-1)/limit + 1; i <= zero; i++ {
		f0[i] = comb(zero-1, i-1)
		for j := 1; j <= (zero-i)/limit; j++ {
			f0[i] = (f0[i] + (1-j%2*2)*comb(i, j)*comb(zero-j*limit-1, i-1)) % mod
		}
	}

	for i := (one-1)/limit + 1; i <= min(one, zero+1); i++ {
		f1 := comb(one-1, i-1)
		for j := 1; j <= (one-i)/limit; j++ {
			f1 = (f1 + (1-j%2*2)*comb(i, j)*comb(one-j*limit-1, i-1)) % mod
		}
		ans = (ans + (f0[i-1]+f0[i]*2+f0[i+1])*f1) % mod
	}
	return (ans + mod) % mod // 保证结果非负
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func numberOfStableArrays2(zero, one, limit int) int {
	const mod = 1_000_000_007
	memo := make([][][2]int, zero+1)
	for i := range memo {
		memo[i] = make([][2]int, one+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) (res int) {
		if i == 0 { // 递归边界
			if k == 1 && j <= limit {
				return 1
			}
			return
		}
		if j == 0 { // 递归边界
			if k == 0 && i <= limit {
				return 1
			}
			return
		}
		p := &memo[i][j][k]
		if *p != -1 { // 之前计算过
			return *p
		}
		if k == 0 {
			// +mod 保证答案非负
			res = (dfs(i-1, j, 0) + dfs(i-1, j, 1)) % mod
			if i > limit {
				res = (res - dfs(i-limit-1, j, 1) + mod) % mod
			}
		} else {
			res = (dfs(i, j-1, 0) + dfs(i, j-1, 1)) % mod
			if j > limit {
				res = (res - dfs(i, j-limit-1, 0) + mod) % mod
			}
		}
		*p = res // 记忆化
		return
	}
	return (dfs(zero, one, 0) + dfs(zero, one, 1)) % mod
}
func main() {
	fmt.Println(numberOfStableArrays(1, 2, 1))
}
