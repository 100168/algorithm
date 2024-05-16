package main

import (
	"fmt"
	"slices"
)

/*
*给你一个整数 n 和一个下标从 0 开始的字符串数组 words ，和一个下标从 0 开始的数组 groups ，两个数组长度都是 n 。

两个长度相等字符串的 汉明距离 定义为对应位置字符 不同 的数目。

你需要从下标 [0, 1, ..., n - 1] 中选出一个 最长子序列，将这个子序列记作长度为 k 的 [i0, i1, ..., ik - 1] ，它需要满足以下条件：

相邻 下标对应的 groups 值 不同。即，对于所有满足 0 < j + 1 < k 的 j 都有 groups[ij] != groups[ij + 1] 。
对于所有 0 < j + 1 < k 的下标 j ，都满足 words[ij] 和 words[ij + 1] 的长度 相等 ，且两个字符串之间的 汉明距离 为 1 。
请你返回一个字符串数组，它是下标子序列 依次 对应 words 数组中的字符串连接形成的字符串数组。如果有多个答案，返回任意一个。

子序列 指的是从原数组中删掉一些（也可能一个也不删掉）元素，剩余元素不改变相对位置得到的新的数组。

注意：words 中的字符串长度可能 不相等 。

输入：n = 3, words = ["bab","dab","cab"], groups = [1,2,2]
输出：["bab","cab"]
解释：一个可行的子序列是 [0,2] 。
- groups[0] != groups[2]
- words[0].length == words[2].length 且它们之间的汉明距离为 1 。
所以一个可行的答案是 [words[0],words[2]] = ["bab","cab"] 。
另一个可行的子序列是 [0,1] 。
- groups[0] != groups[1]
- words[0].length = words[1].length 且它们之间的汉明距离为 1 。
所以另一个可行的答案是 [words[0],words[1]] = ["bab","dab"] 。
符合题意的最长子序列的长度为 2 。
*/
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)
	dp := make([]int, n)

	//需要一个辅助数组记录是由哪个状态转移而来
	from := make([]int, n)
	for i := 0; i < n; i++ {
		curStr := words[i]
		curVal := groups[i]
		dp[i] = 1
		from[i] = i
		for j := 0; j < i; j++ {
			if len(words[j]) != len(curStr) || curVal == groups[j] {
				continue
			}
			diff := 0
			for k := 0; k < len(curStr); k++ {
				if words[j][k] != curStr[k] {
					diff++
				}
			}
			if diff != 1 {
				continue
			}
			dp[i] = max(dp[i], dp[j]+1)
			if dp[i] == dp[j]+1 {
				from[i] = j
			}
		}
	}
	//找到最大值的位置
	maxVal := 0
	maxIndex := 0
	for i, v := range dp {
		if v > maxVal {
			maxVal = v
			maxIndex = i
		}
	}
	ans := make([]string, 0)
	//index相同则跳出
	for from[maxIndex] != maxIndex {
		ans = append(ans, words[maxIndex])
		maxIndex = from[maxIndex]
	}
	//最后将自己也加进去
	ans = append(ans, words[maxIndex])
	//因为是从后往前加，需要反转一下
	slices.Reverse(ans)
	return ans
}

func main() {
	fmt.Println(getWordsInLongestSubsequence([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
}
