package main

import "fmt"

/*
*
给你一个正整数 n。

如果一个二进制字符串 x 的所有长度为 2 的
子字符串
中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。

返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。
*/
func validStrings(n int) []string {

	ans := make([]string, 0)

	ans = append(ans, "1")
	ans = append(ans, "0")

	if n == 1 {
		return []string{"1"}
	}
	for i := 1; i < n; i++ {
		temp := ans
		ans = nil
		for _, v := range temp {
			ans = append(ans, v+string('1'))
			if v[len(v)-1] != '0' {
				ans = append(ans, v+string('0'))
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(validStrings(3))
}
