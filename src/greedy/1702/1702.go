package main

import (
	"fmt"
)

//给你一个二进制字符串 binary ，它仅有 0 或者 1 组成。你可以使用下面的操作任意次对它进行修改：
//
//
// 操作 1 ：如果二进制串包含子字符串 "00" ，你可以用 "10" 将其替换。
//
//
//
// 比方说， "00010" -> "10010"
//
//
// 操作 2 ：如果二进制串包含子字符串 "10" ，你可以用 "01" 将其替换。
//
// 比方说， "00010" -> "00001"
//
//
//
//
// 请你返回执行上述操作任意次以后能得到的 最大二进制字符串 。如果二进制字符串 x 对应的十进制数字大于二进制字符串 y 对应的十进制数字，那么我们称二进制
//字符串 x 大于二进制字符串 y 。
//
//
//
// 示例 1：
//
//
//输入：binary = "000110"
//输出："111011"
//解释：一个可行的转换为：
//"000110" -> "000101"
//"000101" -> "100101"
//"100101" -> "110101"
//"110101" -> "110011"
//"110011" -> "111011"
//
//
// 示例 2：
//
//
//输入：binary = "01"
//输出："01"
//解释："01" 没办法进行任何转换。
//
//
//
//
// 提示：
//
//
// 1 <= binary.length <= 10⁵
// binary 仅包含 '0' 和 '1' 。
//
//
// Related Topics 贪心 字符串 👍 59 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func maximumBinaryString(binary string) string {
	n := len(binary)
	s := []rune(binary)
	j := 0
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			for j <= i || j < n && s[j] == '1' {
				j++
			}
			if j < n {
				s[j] = '1'
				s[i] = '1'
				s[i+1] = '0'
			}
		}
	}
	return string(s)
}

func main() {
	fmt.Println(maximumBinaryString("000110"))
	fmt.Println(maximumBinaryString("10011001"))
	fmt.Println(maximumBinaryString("01010"))
}

//leetcode submit region end(Prohibit modification and deletion)
