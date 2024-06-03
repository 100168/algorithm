package main

import "fmt"

/*
*
给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。
分数可以以 任意 顺序返回。
*/
func simplifiedFractions(n int) []string {

	var ans []string

	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, fmt.Sprintf("%d%s%d", i, "/", j))
			}
		}
	}
	return ans
}
func gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(simplifiedFractions(10))
}
