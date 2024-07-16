package main

import "fmt"

/*
*
给你一个字符串 s 和一个 正 整数 k 。

从字符串 s 中选出一组满足下述条件且 不重叠 的子字符串：

每个子字符串的长度 至少 为 k 。
每个子字符串是一个 回文串 。
返回最优方案中能选择的子字符串的 最大 数目。

子字符串 是字符串中一个连续的字符序列。

思路:
1.回文判断+贪心
*/
func maxPalindromes(s string, k int) int {

	n := len(s)
	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	ans := 0
	right := -1

	if k == 1 {
		return len(s)
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if i-1 == j {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j+1] && s[i] == s[j] {
					dp[i][j] = true
				}
			}
			if dp[i][j] == true && (i-j+1) >= k && j > right {
				ans++
				right = i
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPalindromes("abaccdbbd", 3))
	fmt.Println(maxPalindromes("fttfjofpnpfydwdwdnns", 2))
	fmt.Println(maxPalindromes("iqqibcecvrbxxj", 1))
}
