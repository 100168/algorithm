package main

import (
	"fmt"
)

/*
*
给你一个字符串 s 和一个整数 k。请你使用以下算法加密字符串：

对于字符串 s 中的每个字符 c，用字符串中 c 后面的第 k 个字符替换 c（以循环方式）。
返回加密后的字符串。
*/
func getEncryptedString(s string, k int) string {

	bytes := []byte(s)
	n := len(bytes)

	ans := make([]byte, n)
	for i := 0; i < len(bytes); i++ {
		ans[i] = bytes[(i+k)%n]
	}
	return string(ans)
}

func main() {
	fmt.Println(getEncryptedString("dart", 3))
}
