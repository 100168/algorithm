package main

import "fmt"

/**
给你一个长度为 n 的字符串 s 和一个整数 k ，n 是 k 的 倍数 。你的任务是将字符串 s 哈希为一个长度为 n / k 的新字符串 result 。

首先，将 s 分割成 n / k 个
子字符串
 ，每个子字符串的长度都为 k 。然后，将 result 初始化为一个 空 字符串。

我们依次从前往后处理每一个 子字符串 ：

一个字符的 哈希值 是它在 字母表 中的下标（也就是 'a' → 0 ，'b' → 1 ，... ，'z' → 25）。
将子字符串中字幕的 哈希值 求和。
将和对 26 取余，将结果记为 hashedChar 。
找到小写字母表中 hashedChar 对应的字符。
将该字符添加到 result 的末尾。
返回 result 。
*/

func stringHash(s string, k int) string {

	n := len(s)

	c := 0

	ans := ""
	for i := 0; i < n; i++ {

		cur := int(s[i] - 'a')

		c += cur
		c %= 26

		if (i+1)%k == 0 {
			ans += string(byte(c + 'a'))
			c = 0
		}
	}

	return ans
}

func main() {
	fmt.Println(stringHash("abcd", 2))
}
