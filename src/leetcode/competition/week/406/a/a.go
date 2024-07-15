package main

/*
*
给你一个仅由数字组成的字符串 s，在最多交换一次 相邻 且具有相同 奇偶性 的数字后，返回可以得到的
字典序最小的字符串
。

如果两个数字都是奇数或都是偶数，则它们具有相同的奇偶性。例如，5 和 9、2 和 4 奇偶性相同，而 6 和 9 奇偶性不同。
*/
func getSmallestString(s string) string {

	bytes := []byte(s)
	for i := 1; i < len(s); i++ {

		if s[i]&'1' == s[i-1]&'1' {
			if s[i] < s[i-1] {

				bytes[i], bytes[i-1] = bytes[i-1], bytes[i]
				return string(bytes)
			}
		}
	}
	return s
}
