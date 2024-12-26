package main

import "fmt"

/*
给你两个字符串 word1 和 word2 。

如果一个字符串 x 修改 至多 一个字符会变成 y ，那么我们称它与 y 几乎相等 。

如果一个下标序列 seq 满足以下条件，我们称它是 合法的 ：

下标序列是 升序 的。
将 word1 中这些下标对应的字符 按顺序 连接，得到一个与 word2 几乎相等 的字符串。
请你返回一个长度为 word2.length 的数组，表示一个
字典序最小的 合法 下标序列。如果不存在这样的序列，请你返回一个 空 数组。

注意 ，答案数组必须是字典序最小的下标数组，而 不是 由这些下标连接形成的字符串。

示例 1：

输入：word1 = "vbcca", word2 = "abc"

输出：[0,1,2]

解释：

字典序最小的合法下标序列为 [0, 1, 2] ：

将 word1[0] 变为 'a' 。
word1[1] 已经是 'b' 。
word1[2] 已经是 'c' 。
示例 2：

输入：word1 = "bacdc", word2 = "abc"

输出：[1,2,4]

解释：

字典序最小的合法下标序列为 [1, 2, 4] ：

word1[1] 已经是 'a' 。
将 word1[2] 变为 'b' 。
word1[4] 已经是 'c' 。
示例 3：

输入：word1 = "aaaaaa", word2 = "aaabc"

输出：[]

解释：

没有合法的下标序列。

示例 4：

输入：word1 = "abc", word2 = "ab"

输出：[0,1]

提示：

1 <= word2.length < word1.length <= 3 * 105
word1 和 word2 只包含小写英文字母。
*/
func validSequence(word1 string, word2 string) []int {
	p := 0
	n := len(word1)
	left := make([]int, n)
	for i, v := range word1 {
		left[i] = p
		if p < len(word2) && word2[p] == byte(v) {
			p++
		}
	}
	ans := -1
	p = len(word2) - 1
	for i := n - 1; i >= 0; i-- {
		v := word1[i]
		if left[i]+len(word2)-1-p >= len(word2)-1 {
			ans = i
		}
		if p >= 0 && word2[p] == v {
			p--
		}
	}
	if ans == -1 {
		return []int{}
	}
	p = 0
	res := make([]int, 0)

	cnt := 1
	for i, v := range word1 {
		if i == ans || byte(v) == word2[p] {
			res = append(res, i)

			if i == ans && byte(v) != word2[p] {
				cnt--
			}
			p++
		} else if i > ans && cnt == 1 {
			cnt--
			res = append(res, i)
			p++
		}
		if p >= len(word2) {
			break
		}
	}
	return res
}

func main() {
	//fmt.Println(validSequence("bacdc", "abc"))
	//fmt.Println(validSequence("ccbccccbcc", "b"))
	//fmt.Println(validSequence("ghhgghhhhhh", "gg"))
	fmt.Println(validSequence("abc", "ab"))
}