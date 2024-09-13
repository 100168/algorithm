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

func waysToFillArray(queries [][]int) []int {

	n := len(queries)

	mod := int(1e9 + 7)

	ans := make([]int, n)

	for i, v := range queries {

		c := v[0]
		x := v[1]

		s := 0
		for j := 2; j*j <= x; j++ {
			for x%j == 0 {

				s++
				x /= j
			}
		}
		if x > 1 {
			s++
		}
		cur := 1

		for ; s > 0; s >>= 1 {

			if s&1 != 0 {

				cur = cur * c % mod
			}
			c = c * c % mod
		}
		ans[i] = cur

	}
	return ans
}

func waysToFillArray2(queries [][]int) []int {

	n := len(queries)

	maxV := 10000 + 1
	mod := int(1e9 + 7)

	ans := make([]int, n)

	comb := make([][]int, maxV)

	for i := range comb {
		comb[i] = make([]int, maxV)
	}

	comb[0][0] = 1

	for i := 1; i < 16; i++ {
		comb[i][i] = 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % mod
		}
	}

	p := make([]int, maxV)
	p[0] = 0
	for i := 1; i < maxV; i++ {
		p[i] = p[i-1] * i % mod
	}

	for i, v := range queries {

		c := v[0]
		x := v[1]

		primeCnt := make([]int, 0)

		curCnt := 0
		s := 0
		for j := 2; j*j <= x; j++ {
			for x%j == 0 {
				curCnt++
				x /= j
			}
			s += curCnt
			if curCnt > 0 {
				primeCnt = append(primeCnt, curCnt)
				curCnt = 0
			}
		}
		if x > 0 {
			primeCnt = append(primeCnt, 1)
			s += 1
		}
		cur := 1

		for ; s > 0; s >>= 1 {

			if s&1 == 0 {

				cur = cur * c % mod
			}
			c = c * c % mod
		}
		ans[i] = c

	}
	return ans
}

func main() {
	fmt.Println(waysToFillArray([][]int{{2, 6}, {5, 1}, {73, 660}}))
}
