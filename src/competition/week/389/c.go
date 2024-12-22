package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
给你一个字符串 word 和一个整数 k。

如果 |freq(word[i]) - freq(word[j])| <= k 对于字符串中所有下标 i 和 j  都成立，则认为 word 是 k 特殊字符串。

此处，freq(x) 表示字符 x 在 word 中的出现频率，而 |y| 表示 y 的绝对值。

返回使 word 成为 k 特殊字符串 需要删除的字符的最小数量。

输入：word = "aabcaba", k = 0

输出：3
解释：可以删除 2 个 "a" 和 1 个 "c" 使 word 成为 0 特殊字符串。word 变为 "baba"，

此时 freq('a') == freq('b') == 2。
*/
func minimumDeletions(word string, k int) int {

	cnt := make([]int, 26)

	for _, v := range word {
		cnt[v-'a']++
	}

	ans := math.MaxInt
	sort.Ints(cnt)

	s := 0
	for i, v := range cnt {

		if v == 0 {
			continue
		}

		cur := s

		for j := 25; j > i; j-- {
			if v+k < cnt[j] {
				cur += cnt[j] - v - k
			} else {
				break
			}
		}
		ans = min(ans, cur)
		s += v
	}
	return ans

}

func main() {
	//fmt.Println(minimumDeletions("aamcaba", 0))
	fmt.Println(minimumDeletions("aabcaba", 0))
}
