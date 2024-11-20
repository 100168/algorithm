package main

import "fmt"

/*
*
给你一个下标从 0 开始的字符串 word 和一个整数 k 。

在每一秒，你必须执行以下操作：

移除 word 的前 k 个字符。
在 word 的末尾添加 k 个任意字符。
注意 添加的字符不必和移除的字符相同。但是，必须在每一秒钟都执行 两种 操作。

返回将 word 恢复到其 初始 状态所需的 最短 时间（该时间必须大于零）。



示例 1：

输入：word = "abacaba", k = 3
输出：2
解释：
第 1 秒，移除 word 的前缀 "aba"，并在末尾添加 "bac" 。因此，word 变为 "cababac"。
第 2 秒，移除 word 的前缀 "cab"，并在末尾添加 "aba" 。因此，word 变为 "abacaba" 并恢复到始状态。
可以证明，2 秒是 word 恢复到其初始状态所需的最短时间。
*/

func minimumTimeToInitialState(word string, k int) int {

	n := len(word)

	z := make([]int, n)

	for i, c, r := 1, 1, 1; i < n; i++ {
		l := 0
		if i < r {
			l = min(r-i, z[i-c])
		}
		for i+l < n && word[l] == word[i+l] {
			l++
		}

		if i+l == n && i%k == 0 {
			return i / k
		}
		if i+l > r {
			c = i
			r = i + l
		}
		z[i] = l
	}
	return (n-1)/k + 1
}

func main() {
	fmt.Println(minimumTimeToInitialState("aa", 1))

	fmt.Println(500 * 500)
}
