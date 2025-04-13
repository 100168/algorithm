package main

import (
	"fmt"
	"slices"
)

/*
*
给你两个字符串 s 和 t。

你可以从 s 中选择一个子串（可以为空）以及从 t 中选择一个子串（可以为空），然后将它们 按顺序 连接，得到一个新的字符串。

返回可以由上述方法构造出的 最长 回文串的长度。

回文串 是指正着读和反着读都相同的字符串。

子字符串 是指字符串中的一个连续字符序列。

思路：最长公共子串+中心扩散or manachar
*/
func getMax(s string, t string) int {

	m := len(s)
	n := len(t)

	t = reverse(t)

	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	fm := make([]int, m+1)
	for i := 1; i <= m; i++ {

		mx := 0
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				f[i][j] = f[i-1][j-1] + 1
				mx = max(mx, f[i][j], mx)
			}
		}
		fm[i] = mx
	}

	ans := slices.Max(fm) * 2

	p := manachar(s)

	for i, v := range p {

		ans = max(ans, v-1+fm[(i-v+1)/2]*2)
	}

	return ans
}

func manachar(s string) []int {

	bytes := make([]byte, len(s)*2+1)

	for i := range bytes {

		if i&1 != 0 {
			bytes[i] = s[i/2]
		} else {
			bytes[i] = '#'
		}
	}

	p := make([]int, len(bytes))

	for i, r, c := 0, 0, 0; i < len(bytes); i++ {

		l := 0
		if r > i {
			l = min(p[2*c-i], r-i)
		}

		for l+i < len(bytes) && i-l >= 0 && bytes[l+i] == bytes[i-l] {
			l++
		}

		if i+l > r {
			r = i + l
			c = i
		}
		p[i] = l

	}
	return p
}

func reverse(t string) string {

	r := ""

	for i := 0; i < len(t); i++ {

		r = string(t[i]) + r

	}
	return r
}

func longestPalindrome(a, b string) int {
	return max(getMax(a, b), getMax(reverse(b), reverse(a)))
}
func main() {

	fmt.Println(longestPalindrome("abcde", "ecdba"))
	fmt.Println(longestPalindrome("b", "aaa"))
	fmt.Println(longestPalindrome("a", "a"))
}
