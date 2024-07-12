package main

import (
	"fmt"
	"unicode"
)

/**
给定一个编码字符串 s 。请你找出 解码字符串 并将其写入磁带。解码时，从编码字符串中 每次读取一个字符 ，并采取以下步骤：

如果所读的字符是字母，则将该字母写在磁带上。
如果所读的字符是数字（例如 d），则整个当前磁带总共会被重复写 d-1 次。
现在，对于给定的编码字符串 s 和索引 k，查找并返回解码字符串中的第 k 个字母。

输入：s = "leet2code3", k = 10
输出："o"
解释：
解码后的字符串为 "leetleetcode leetleetcode leetleetcode"。
字符串中的第 10 个字母是 "o"。

思路:
1.找到>=k 的长度 并记录下标
2.然后

*/

func decodeAtIndex(s string, k int) string {

	l := 0
	index := 0
	for i, v := range s {

		if unicode.IsLetter(v) {
			l++
		} else {
			l *= int(s[i] - '0')
		}
		if l >= k {
			index = i
			break
		}
	}
	for {
		if unicode.IsLetter(rune(s[index])) {
			if l == k {
				return string(s[index])
			}
			l--
		} else {
			l /= int(s[index] - '0')
			k %= l
			if k == 0 {
				k = l
			}
		}
		index--
	}

}

func main() {
	fmt.Println(decodeAtIndex("leet2code3", 10))
	fmt.Println(decodeAtIndex("habbc4", 6))
	fmt.Println(decodeAtIndex("ab234567d9d9e9e99e99f9m9", 1<<38-100))
}
