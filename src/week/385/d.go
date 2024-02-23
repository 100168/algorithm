package main

/*
100208. 统计前后缀下标对 II 显示英文描述
通过的用户数277
尝试过的用户数915
用户总通过次数426
用户总提交次数2963
题目难度Hard
给你一个下标从 0 开始的字符串数组 words 。

定义一个 布尔 函数 isPrefixAndSuffix ，它接受两个字符串参数 str1 和 str2 ：

当 str1 同时是 str2 的前缀（prefix）和后缀（suffix）时，isPrefixAndSuffix(str1, str2) 返回 true，否则返回 false。
例如，isPrefixAndSuffix("aba", "ababa") 返回 true，因为 "aba" 既是 "ababa" 的前缀，也是 "ababa" 的后缀，但是 isPrefixAndSuffix("abc", "abcd") 返回 false。

以整数形式，返回满足 i < j 且 isPrefixAndSuffix(words[i], words[j]) 为 true 的下标对 (i, j) 的 数量 。*/

func countPrefixSuffixPairs2(words []string) int64 {

	ans := int64(0)

	prefixMap := make(map[string]int)

	n := len(words)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && words[z[i]] == words[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	for _, v := range words {
		for i := 0; i < len(v); i++ {
			if v[:i+1] == v[len(v)-i-1:] {
				prefixMap[v[:i+1]]++
			}
		}
	}

	for _, v := range words {
		for i := 0; i < len(v); i++ {
			if v[:i+1] == v[len(v)-i-1:] {
				prefixMap[v[:i+1]]--
			}
			ans += int64(prefixMap[v])
		}
	}

	return ans

}
func main() {

	println(countPrefixSuffixPairs2([]string{"aab", "a", "cba", "a"}))
}
