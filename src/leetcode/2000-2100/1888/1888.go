package main

import (
	"fmt"
	"math"
	"math/rand"
)

/**
给你一个二进制字符串 s 。你可以按任意顺序执行以下两种操作任意次：

类型 1 ：删除 字符串 s 的第一个字符并将它 添加 到字符串结尾。
类型 2 ：选择 字符串 s 中任意一个字符并将该字符 反转 ，也就是如果值为 '0' ，则反转得到 '1' ，反之亦然。
请你返回使 s 变成 交替 字符串的前提下， 类型 2 的 最少 操作次数 。

我们称一个字符串是 交替 的，需要满足任意相邻字符都不同。
比方说，字符串 "010" 和 "1010" 都是交替的，但是字符串 "0100" 不是。

1010101010

0000100011

思路：前后缀分解

我是废物

*/

func minFlips(s string) int {
	return min(solve(s, '1'), solve(s, '0'))
}

func solve(s string, head byte) int {

	n := len(s)

	left := make([]int, n)
	diff := 0
	for i := 0; i < n; i++ {
		//这个有点离谱
		if s[i] != head^byte(i&1) {
			diff++
		}
		left[i] = diff
	}
	ans := math.MaxInt

	diff = 0
	tail := head ^ 1
	for i := n - 1; i >= 0; i-- {
		ans = min(ans, left[i]+diff)
		if s[i] != tail^byte((n-i+1)&1) {
			diff++
		}
	}
	return ans
}

func genTest(n int) string {

	ans := ""

	for i := 0; i < n; i++ {
		seed := rand.Intn(2)
		if seed == 0 {
			ans += "0"
		} else {
			ans += "1"
		}

	}
	return ans
}

func main() {

	//fmt.Println(genTest(10))
	fmt.Println(minFlips("0000100011"))
}
