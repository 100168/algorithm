package main

import "strconv"

//给你一个整数 n ，以二进制字符串的形式返回该整数的 负二进制（base -2）表示。
//
// 注意，除非字符串就是 "0"，否则返回的字符串中不能含有前导零。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出："110"     010
//解释：(-2)² + (-2)¹ = 2
//
//
// 示例 2：
//
//
//输入：n = 3,
//输出："111"  011
//解释：(-2)² + (-2)¹ + (-2)⁰ = 4-2+1=3
//
//
// 示例 3：
//
//
//输入：n = 4
//输出："100"  100
//解释：(-2)² = 4
// 11    1011 = (-2)3 + (-2)1+(-2)0 = -8-2+1 + 16
//
//
//
// 提示：
//
//
// 0 <= n <= 10⁹
//1010
//
//

func baseNeg2(n int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}
	var res []byte
	for n != 0 {
		remainder := n & 1
		res = append(res, '0'+byte(remainder))
		n -= remainder
		n /= -2
	}
	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return string(res)
}
