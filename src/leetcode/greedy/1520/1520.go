package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/**
给你一个只包含小写字母的字符串 s ，你需要找到 s 中最多数目的非空子字符串，满足如下条件：

这些字符串之间互不重叠，也就是说对于任意两个子字符串 s[i..j] 和 s[x..y] ，要么 j < x 要么 i > y 。
如果一个子字符串包含字符 char ，那么 s 中所有 char 字符都应该在这个子字符串中。
请你找到满足上述条件的最多子字符串数目。如果有多个解法有相同的子字符串数目，请返回这些子字符串总长度最小的一个解。可以证明最小总长度解是唯一的。

请注意，你可以以 任意 顺序返回最优解的子字符串。



示例 1：

输入：s = "adefaddaccc"
输出：["e","f","ccc"]
解释：下面为所有满足第二个条件的子字符串：
[
  "adefaddaccc"
  "adefadda",
  "ef",
  "e",
  "f",
  "ccc",
]
如果我们选择第一个字符串，那么我们无法再选择其他任何字符串，所以答案为 1 。
如果我们选择 "adefadda" ，剩下子字符串中我们只可以选择 "ccc" ，
它是唯一不重叠的子字符串，所以答案为 2 。同时我们可以发现，选择 "ef" 不是最优的，因为它可以被拆分成 2 个子字符串。所以最优解是选择 ["e","f","ccc"] ，答案为 3 。不存在别的相同数目子字符串解。
示例 2：

输入：s = "abbaccd"
输出：["d","bb","cc"]
解释：注意到解 ["d","abba","cc"] 答案也为 3 ，但它不是最优解，因为它的总长度更长。


提示：

1 <= s.length <= 10^5
s 只包含小写英文字母。


md, 题都看不懂
*/

func maxNumOfSubstrings(s string) []string {

	indexMap := make(map[rune][]int)
	for i := 'a'; i <= 'z'; i++ {
		first := strings.Index(s, string(i))
		last := strings.LastIndex(s, string(i))
		if first == -1 {
			continue
		}
		indexMap[i] = []int{first, last}
	}

	newPair := make([][]int, 0)
	for _, v := range indexMap {

		newPair = append(newPair, v)

	}
	sort.Slice(newPair, func(i, j int) bool {
		return newPair[i][0] < newPair[j][0]
	})
	for _, v := range newPair {
		k := s[v[0]]

		if _, ok := indexMap[rune(k)]; !ok {
			continue
		}
		for i := v[0]; i <= v[1]; i++ {
			cur := indexMap[rune(s[i])]
			v[0] = min(v[0], cur[0])
			v[1] = max(cur[1], v[1])
		}
		indexMap[rune(k)] = v
	}

	newPair = make([][]int, 0)
	for _, v := range indexMap {

		newPair = append(newPair, v)

	}

	sort.Slice(newPair, func(i, j int) bool {
		return newPair[i][1] < newPair[j][1]
	})

	end := math.MinInt

	var ans []string
	for _, v := range newPair {
		if v[0] > end {
			end = v[1]
			ans = append(ans, s[v[0]:v[1]+1])
		}
	}
	return ans

}

func main() {
	fmt.Println(maxNumOfSubstrings("adefaddaccc"))
	fmt.Println(maxNumOfSubstrings("abbaccd"))
	fmt.Println(maxNumOfSubstrings("abab"))
	fmt.Println(maxNumOfSubstrings("ababa"))
	fmt.Println(maxNumOfSubstrings("abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(maxNumOfSubstrings("cabcccbaa"))
	fmt.Println(maxNumOfSubstrings("bbcacbaba"))
}
