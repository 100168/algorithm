package main

import "fmt"

/*
*
给你一维空间的 n 个点，其中第 i 个点（编号从 0 到 n-1）位于 x = i 处，
请你找到 恰好 k 个不重叠 线段且每个线段至少覆盖两个点的方案数。线段的两个端点必须都是 整数坐标 。
这 k 个线段不需要全部覆盖全部 n 个点，且它们的端点 可以 重合。

请你返回 k 个不重叠线段的方案数。由于答案可能很大，请将结果对 109 + 7 取余 后返回。
*/
func numberOfSets(n int, k int) int {

	mod := int(1e9 + 7)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)

		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][0][0] = 1
	//把点转成线段考虑然后再dp
	for i := 1; i < n; i++ {
		for j := 0; j <= k && j <= i; j++ {
			dp[i][j][0] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
			dp[i][j][1] = dp[i-1][j][1]
			if j > 0 {
				dp[i][j][1] = (dp[i][j][1] + dp[i-1][j-1][1] + dp[i-1][j-1][0]) % mod
			}

		}
	}
	return (dp[n-1][k][0] + dp[n-1][k][1]) % mod
}

/*
n 个点分隔出 n-1 条单位线段。

该题等价于：把 n-1 条单位线段切分成 2k+1 份，其中第 2, 4,..., 2k 共 k 份的总长度 ≥1 ，其余 2k+1 份的长度可为零，求分割方法总数。

解决该问题可以使用高中排列组合时学过的“隔板法”。

c(n-1+n-1-k)(n-1-k)

c(n-1+k)(k-1)
*/

const mod = int(1e9 + 7)

func numberOfSets2(n int, k int) int {

	n = n - 1 + k
	ans := n

	//2*n-2+k
	//n-1-k
	div := 1
	for i := 2; i <= k*2; i++ {
		ans = ans * (n - i + 1) % mod
		div = div * i % mod
	}
	return ans * power(div, mod-2) % mod
}

func power(x, y int) int {
	ans := 1
	for ; y > 0; y >>= 1 {
		if y&1 > 0 {
			ans = ans * x % mod
		}
		x = x * x % mod
	}
	return ans
}

func main() {
	fmt.Println(numberOfSets(10, 7))
	fmt.Println(numberOfSets2(10, 7))
}
