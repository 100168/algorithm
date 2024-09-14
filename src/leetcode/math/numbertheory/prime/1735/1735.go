package main

import "fmt"

/**
给你一个二维整数数组 queries ，其中 queries[i] = [ni, ki] 。
第 i 个查询 queries[i] 要求构造长度为 ni 、每个元素都是正整数的数组，且满足所有元素的乘积为 ki ，请你找出有多少种可行的方案。
由于答案可能会很大，方案数需要对 109 + 7 取余 。

请你返回一个整数数组 answer，满足 answer.length == queries.length ，其中 answer[i]是第 i 个查询的结果。



示例 1：

输入：queries = [[2,6],[5,1],[73,660]]
输出：[4,1,50734910]
解释：每个查询之间彼此独立。
[2,6]：总共有 4 种方案得到长度为 2 且乘积为 6 的数组：[1,6]，[2,3]，[3,2]，[6,1]。
[5,1]：总共有 1 种方案得到长度为 5 且乘积为 1 的数组：[1,1,1,1,1]。
[73,660]：总共有 1050734917 种方案得到长度为 73 且乘积为 660 的数组。1050734917 对 109 + 7 取余得到 50734910 。
示例 2 ：

输入：queries = [[1,1],[2,2],[3,3],[4,4],[5,5]]
输出：[1,2,3,10,5]

2,   2,2,3    分成 2组

1,cnt2,cnt3

cn1,cn2,cn3



caod

*/

const mod = int(1e9 + 7)

var p [10015]int

const k = 10015

func init() {
	p[0] = 1
	for i := 1; i < k; i++ {
		p[i] = p[i-1] * i % mod
	}
}

func comb(n, m int) int {
	return p[n] * getRp(p[m]*p[n-m]%mod, mod) % mod
}

func waysToFillArray(queries [][]int) []int {

	l := len(queries)

	ans := make([]int, l)

	for i, v := range queries {
		c := v[0]
		x := v[1]

		cur := 1

		for j := 2; j*j <= x; j++ {
			curCnt := 0
			for x%j == 0 {
				curCnt++
				x /= j
			}

			if curCnt > 0 {
				//总个数
				cur = cur * comb(c+curCnt-1, c-1) % mod
			}
		}
		if x > 1 {
			cur = cur * comb(c, c-1) % mod
		}
		ans[i] = cur
	}
	return ans
}

func pow(a, b int) int {
	ans := 1
	for ; b > 0; b >>= 1 {
		if b&1 != 0 {
			ans = ans * a % mod
		}
		a = a * a % mod
	}
	return ans
}

func exgGcd(a, b int) (int, int, int) {
	if b == 0 {
		return a, 0, a
	} else {
		x1, y1, d := exgGcd(b, a%b)
		x := y1
		y := x1 - y1*(a/b)
		return x, y, d
	}
}

func getRp(a, b int) int {

	x, _, _ := exgGcd(a, b)
	return (x + mod) % mod
}
func main() {
	fmt.Println(waysToFillArray([][]int{{2, 6}, {5, 1}, {73, 660}}))

	x, _, _ := exgGcd(1024305, mod)
	fmt.Println((x + mod) % mod)
	fmt.Println(pow(1024305, mod-2))
}
