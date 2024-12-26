package main

import (
	"fmt"
	"math"
)

/*
*
给你一个整数数组 nums。

请你统计所有满足一下条件的 非空
子序列对 (seq1, seq2) 的数量：

子序列 seq1 和 seq2 不相交，意味着 nums 中 不存在 同时出现在两个序列中的下标。
seq1 元素的GCD等于 seq2 元素的 GCD。

返回满足条件的子序列对的总数。

由于答案可能非常大，请返回其对 109 + 7 取余 的结果。

示例 1：

输入： nums = [1,2,3,4]

输出： 10

解释：

元素 GCD 等于 1 的子序列对有：

示例 2：

输入： nums = [10,20,30]

输出： 2

解释：

元素 GCD 等于 10 的子序列对有：

([10, 20, 30], [10, 20, 30])
([10, 20, 30], [10, 20, 30])
示例 3：

输入： nums = [1,1,1,1]

输出： 50
*/

const mx = 201

var c [mx][mx]int

func init() {

	c[0][0] = 1
	for i := 1; i < mx; i++ {
		c[i][0] = 1
		for j := 1; j <= i; j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}
}

const mod = int(1e9 + 7)

func subsequencePairCount(nums []int) int {

	n := len(nums)

	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, 201)
		for j := range f[i] {
			f[i][j] = make([]int, 201)

			for k := range f[i][j] {
				f[i][j][k] = -1
			}
		}
	}

	var dfs func(int, int, int) int

	dfs = func(i int, l int, r int) int {

		if i < 0 {

			if l == r {
				return 1
			}
			return 0
		}

		if f[i][l][r] != -1 {
			return f[i][l][r]
		}

		cur := (dfs(i-1, l, r) + dfs(i-1, gcd(l, nums[i]), r) + dfs(i-1, l, gcd(r, nums[i]))) % mod
		f[i][l][r] = cur

		return cur
	}

	return dfs(n-1, 0, 0) - 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func subsequencePairCount2(nums []int) int {
	cnt := make([]int, mx+1)
	for _, v := range nums {
		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				cnt[i]++
				if v/i != i {
					cnt[v/i]++
				}
			}
		}
	}
	ans := 0
	//cn1*(cn1(n-1)~c(n-1)(n-1))
	n := len(nums)
	for i := 1; i < n; i++ {
		cur := 0
		for j := 1; j <= n-i; j++ {
			cur = (cur + c[n][i]*c[n-i][j]%mod) % mod
		}
		ans = (ans + cur) % mod
	}

	for _, v := range cnt[2:] {
		for i := 1; i < v; i++ {
			cur := 0
			for j := 1; j <= v-i; j++ {
				cur = (cur + c[v][i]*c[v-i][j]%mod) % mod
			}
			ans = (ans - cur + mod) % mod
		}
	}
	return ans

}

func main() {
	fmt.Println(subsequencePairCount2([]int{1, 2, 3, 4, 5}))

	fmt.Println(math.Round(2.5))
}
