package main

import "fmt"

/*
*给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。

字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。
*/
func countVowelStrings(n int) int {

	ans := 0

	dp := make([]int, 5)

	for i := range dp {
		dp[i] = 1
	}
	re := string("aeiou")
	for i := 1; i < n; i++ {
		temp := make([]int, 5)
		for j := range re {
			for k := 0; k <= j; k++ {
				temp[j] += dp[k]
			}
		}
		dp = temp
	}
	for _, v := range dp {
		ans += v
	}
	return ans
}

// 将n个球放入5个盒子中，盒子可以为空 c(n+4,4)
func countVowelStrings2(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / (1 * 2 * 3 * 4)
}

func main() {
	fmt.Println(countVowelStrings(50))
	fmt.Println(countVowelStrings2(50))
}
