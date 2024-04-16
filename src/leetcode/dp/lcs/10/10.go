package main

//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
// 示例 1：
//
//
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
//
//
// 示例 2:
//
//
//输入：s = "aa", p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
//
// 示例 3：
//
//
//输入：s = "ab", p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// 1 <= p.length <= 20
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
//
// Related Topics 递归 字符串 动态规划 👍 3882 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {

	m := len(s)
	n := len(p)

	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	var dfs func(int, int) bool
	dfs = func(i int, j int) bool {
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 {
			for c := j; c >= 0; c-- {
				if p[c] == '*' {
					c--
				} else {
					return false
				}

			}
			return true
		}
		if j < 0 {
			return false
		}
		if cache[i][j] != -1 {
			return cache[i][j] > 0
		}
		cur := false
		cache[i][j] = 0
		if s[i] == p[j] || p[j] == '.' {
			cur = dfs(i-1, j-1)
		} else if p[j] == '*' {
			for k := i; k >= 0; k-- {
				if s[k] == p[j-1] || p[j-1] == '.' {
					cur = cur || dfs(k-1, j-2)
				} else {
					break
				}
			}
			cur = cur || dfs(i, j-2)
		}
		if cur {
			cache[i][j] = 1
		}
		return cur
	}
	return dfs(m-1, n-1)
}
func main() {
	println(isMatch("aab", "a*b*"))
	println(isMatch("aab", "c*a*b"))
	println(isMatch("aaabaaaababcbccbaab", "c*c*.*c*a*..*c*"))
}

//leetcode submit region end(Prohibit modification and deletion)
