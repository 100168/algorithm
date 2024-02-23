package main

/*100212. 统计前后缀下标对 I 显示英文描述
通过的用户数1851
尝试过的用户数1883
用户总通过次数2103
用户总提交次数3037
题目难度Easy
给你一个下标从 0 开始的字符串数组 words 。

定义一个 布尔 函数 isPrefixAndSuffix ，它接受两个字符串参数 str1 和 str2 ：

当 str1 同时是 str2 的前缀（prefix）和后缀（suffix）时，isPrefixAndSuffix(str1, str2) 返回 true，否则返回 false。
例如，isPrefixAndSuffix("aba", "ababa") 返回 true，因为 "aba" 既是 "ababa" 的前缀，也是 "ababa" 的后缀，但是 isPrefixAndSuffix("abc", "abcd") 返回 false。

以整数形式，返回满足 i < j 且 isPrefixAndSuffix(words[i], words[j]) 为 true 的下标对 (i, j) 的 数量 。*/

func countPrefixSuffixPairs(words []string) int {

	ans := 0
	n := len(words)
	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {
			curI := words[i]
			curJ := words[j]
			if len(curI) > len(curJ) {
				continue
			}
			if curI == curJ[:len(words[i])] && curI == curJ[len(curJ)-len(curI):] {
				ans++
			}

		}
	}
	return ans
}

func main() {
	println(countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa", "a", "aba", "ababa", "aa"}))
}
