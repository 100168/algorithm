package main

import "fmt"

/*
*
Alice 正在她的电脑上输入一个字符串。但是她打字技术比较笨拙，她可能在一个按键上按太久，导致一个字符被输入多次。

给你一个字符串word，它表示最终显示在 Alice 显示屏上的结果。同时给你一个正整数k，表示一开始 Alice 输入字符串的长度至少为k。

请你返回 Alice 一开始可能想要输入字符串的总方案数。

由于答案可能很大，请你将它对109 + 7取余后返回。

示例 1：

输入：word = "aabbccdd", k = 7

输出：5

解释：

可能的字符串包括："aabbccdd"，"aabbccd"，"aabbcdd"，"aabccdd"和"abbccdd"。

示例 2：

输入：word = "aabbccdd", k = 8

输出：1

解释：

唯一可能的字符串是"aabbccdd"。

示例 3：

输入：word = "aaabbb", k = 3

输出：8

提示：

1 <= word.length <= 5 * 105
word只包含小写英文字母。
1 <= k <= 2000

思路：分组背包+前缀和优化dp

反思:为什么做这么慢？

1.总是在下标处理上模糊不清
*/
func possibleStringCount(word string, k int) int {

	mod := 1_000_000_007
	if len(word) < k {
		return 0
	}
	cnt := make([]int, 0)

	sum := 1
	n := len(word)
	for i := 0; i < n; i++ {
		c := 0
		for j := i; j < n && word[j] == word[i]; j++ {
			c++
		}
		if c > 1 {
			cnt = append(cnt, c-1)
		}
		i = i + c - 1
		sum = sum * c % mod
		k--
	}
	if k <= 0 {
		return sum
	}

	m := len(cnt)
	f := make([][]int, m)

	for i := range f {
		f[i] = make([]int, k+1)
	}

	//初始数据
	s := make([]int, k+1)
	for j := 0; j <= min(cnt[0], k); j++ {
		f[0][j] = 1
	}

	for i := 0; i < k; i++ {
		s[i+1] = s[i] + f[0][i]
	}

	//前i组总共长度为j
	//当前组可取范围[0~k]

	for i := 1; i < m; i++ {
		v := cnt[i]
		newS := make([]int, k+1)
		for j := 1; j <= k; j++ {
			f[i][j] = (s[j] - s[max(j-v-1, 0)] + mod) % mod
			newS[j] = (newS[j-1] + f[i][j]) % mod
		}

		s = newS
	}

	return (sum - s[k] + mod) % mod

	//最少留k个
	//  f[i][k]   s = f[i-1][k-1]+f[i-1][k-c]

}

//1  5  11

func main() {
	//fmt.Println(possibleStringCount("aabbccdd", 7))
	fmt.Println(possibleStringCount("aaabbb", 3))
}
