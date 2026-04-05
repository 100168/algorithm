package main

import "strconv"

/*给你两个 正整数 数组 arr1 和 arr2 。

正整数的 前缀 是其 最左边 的一位或多位数字组成的整数。例如，123 是整数 12345 的前缀，而 234 不是 。

设若整数 c 是整数 a 和 b 的 公共前缀 ，那么 c 需要同时是 a 和 b 的前缀。例如，5655359 和 56554 有公共前缀 565 ，而 1223 和 43456 没有 公共前缀。

你需要找出属于 arr1 的整数 x 和属于 arr2 的整数 y 组成的所有数对 (x, y) 之中最长的公共前缀的长度。

返回所有数对之中最长公共前缀的长度。如果它们之间不存在公共前缀，则返回 0 。*/

func longestCommonPrefix(arr1 []int, arr2 []int) int {

	mapA := make(map[string]bool)
	mapB := make(map[string]bool)
	for _, e := range arr1 {
		itoa := strconv.Itoa(e)
		for i := 1; i <= len(itoa); i++ {
			mapA[itoa[:i]] = true
		}
	}
	for _, e := range arr2 {
		itoa := strconv.Itoa(e)
		for i := 1; i <= len(itoa); i++ {
			mapB[itoa[:i]] = true
		}
	}
	ans := 0
	for s := range mapA {
		if mapB[s] {
			ans = max(ans, len(s))
		}
	}
	return ans
}
