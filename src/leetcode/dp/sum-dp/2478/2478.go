package main

import (
	"fmt"
	"strings"
)

/*
*
给你一个字符串 s ，每个字符是数字 '1' 到 '9' ，再给你两个整数 k 和 minLength 。

如果对 s 的分割满足以下条件，那么我们认为它是一个 完美 分割：

s 被分成 k 段互不相交的子字符串。
每个子字符串长度都 至少 为 minLength 。
每个子字符串的第一个字符都是一个 质数 数字，最后一个字符都是一个 非质数 数字。质数数字为 '2' ，'3' ，'5' 和 '7' ，剩下的都是非质数数字。
请你返回 s 的 完美 分割数目。由于答案可能很大，请返回答案对 109 + 7 取余 后的结果。

一个 子字符串 是字符串中一段连续字符串序列。

示例 1：

输入：s = "23542185131", k = 3, minLength = 2
输出：3
解释：存在 3 种完美分割方案：
"2354 | 218 | 5131"
"2354 | 21851 | 31"
"2354218 | 51 | 31"
示例 2：

输入：s = "23542185131", k = 3, minLength = 3
输出：1
解释：存在一种完美分割方案："2354 | 218 | 5131" 。
示例 3：

输入：s = "3312958", k = 3, minLength = 1
输出：1
解释：存在一种完美分割方案："331 | 29 | 58" 。

提示：

1 <= k, minLength <= s.length <= 1000
s 每个字符都为数字 '1' 到 '9' 之一。
*/
func beautifulPartitions(s string, k int, minLength int) int {

	n := len(s)
	mod := int(1e9 + 7)

	sum := make([][]int, n+1)

	f := make([][]int, n+1)

	for i := range sum {
		sum[i] = make([]int, k+1)
		f[i] = make([]int, k+1)
	}

	isPrime := func(c byte) bool {
		return strings.IndexByte("2357", c) >= 0
	}
	if !isPrime(s[0]) || isPrime(s[n-1]) || minLength*k > n {
		return 0
	}
	sum[0][0] = 1
	t := 1
	//  可以把i看成长度
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			sum[i][j] = sum[i-1][j]
		}
		if i >= minLength && !isPrime(s[i-1]) && (i == n || isPrime(s[i])) {
			for j := 1; j <= k; j++ {
				t = sum[i-minLength][j-1]
				sum[i][j] = (sum[i-1][j] + t) % mod
			}
		}
	}
	return t
}

func beautifulPartitions2(s string, k, l int) (ans int) {
	const mod int = 1e9 + 7
	isPrime := func(c byte) bool { return strings.IndexByte("2357", c) >= 0 }
	n := len(s)
	if k*l > n || !isPrime(s[0]) || isPrime(s[n-1]) { // 剪枝
		return
	}
	// 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
	canPartition := func(j int) bool { return j == 0 || j == n || !isPrime(s[j-1]) && isPrime(s[j]) }
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i := 1; i <= k; i++ {
		sum := 0
		// 优化：枚举的起点和终点需要给前后的子串预留出足够的长度
		for j := i * l; j+(k-i)*l <= n; j++ {
			if canPartition(j - l) { // j'=j-l 双指针
				sum = (sum + f[i-1][j-l]) % mod
			}
			if canPartition(j) {
				f[i][j] = sum
			}
		}
	}
	return f[k][n]
}

func main() {
	fmt.Println(beautifulPartitions("23542185131", 3, 2))
}
