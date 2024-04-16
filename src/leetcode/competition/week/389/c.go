package main

import (
	"fmt"
	"math"
)

/*
*
给你一个字符串 word 和一个整数 k。

如果 |freq(word[i]) - freq(word[j])| <= k 对于字符串中所有下标 i 和 j  都成立，则认为 word 是 k 特殊字符串。

此处，freq(x) 表示字符 x 在 word 中的出现频率，而 |y| 表示 y 的绝对值。

返回使 word 成为 k 特殊字符串 需要删除的字符的最小数量。
*/
func minimumDeletions(word string, k int) int {

	cntMap := make(map[uint8]int)

	n := len(word)
	for i := range word {
		cur := word[i] - 'a'
		cntMap[cur]++
	}
	cnt := make([]int, 1)

	for _, v := range cntMap {
		cnt = append(cnt, v)
	}

	ans := math.MaxInt
	for i := 1; i <= n; i++ {
		sb := 0
		for j := 0; j < len(cnt); j++ {
			if cnt[j] < i {
				sb += cnt[j]
			} else {
				sb += max(cnt[j]-i-k, 0)
			}
		}
		ans = min(ans, sb)
	}

	return ans

}

func main() {
	minimum := minimumDeletions("aamcaba", 0)
	fmt.Println(minimum)
}
