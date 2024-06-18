package main

import "fmt"

/*
*
给你一个整数 n 。

如果一个字符串 s 只包含小写英文字母，且 将 s 的字符重新排列后，新字符串包含 子字符串 "leet" ，那么我们称字符串 s 是一个 好 字符串。

比方说：

字符串 "lteer" 是好字符串，因为重新排列后可以得到 "leetr" 。
"letl" 不是好字符串，因为无法重新排列并得到子字符串 "leet" 。
请你返回长度为 n 的好字符串 总 数目。

由于答案可能很大，将答案对 109 + 7 取余 后返回。

子字符串 是一个字符串中一段连续的字符序列。
思路：
1.正南则反
2.先计算所有情况
3.分类讨论计算所有不满足的情况
4.将所有情况-不满足的情况=满足情况

分类讨论：
1.l和t,e都不选或者e只选一次
2.l,t,e选两个不选或者选一个e和并且l,t选一个
3.l,t,e都不选，或者选一个e其他两个不选
4.容斥原理+快速幂
*/

const mod = int(1e9 + 7)

func stringCount(n int) int {
	if n < 4 {
		return 0
	}
	cc := pow(26, n)

	s := (3*pow(25, n) + n*pow(25, n-1) - (3*pow(24, n) + 2*n*pow(24, n-1)) + pow(23, n) + n*pow(23, n-1) + mod) % mod
	return (cc - s + mod) % mod

}

func pow(x, b int) int {
	ans := 1
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			ans = ans * x % mod
		}
		x = x * x % mod
	}
	return ans
}

func stringCount2(n int) int {

	var dfs func(int, int, int, int) int

	memo := make([][][][]int, n+1)
	for i := range memo {
		memo[i] = make([][][]int, 2)
		for j := range memo[i] {
			memo[i][j] = make([][]int, 2)
			for k := range memo[i][j] {
				memo[i][j][k] = make([]int, 3)
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}
	dfs = func(i, l, t, e int) int {
		if i == 0 {
			if l == 0 && t == 0 && e == 0 {
				return 1
			}
			return 0
		}

		if memo[i][l][t][e] != -1 {
			return memo[i][l][t][e]
		}
		cur := 0
		cur = dfs(i-1, 0, t, e)
		cur += dfs(i-1, l, 0, e)
		cur += dfs(i-1, l, t, max(e-1, 0))
		cur += dfs(i-1, l, t, e) * 23
		cur %= mod
		memo[i][l][t][e] = cur
		return cur
	}
	return dfs(n, 1, 1, 2)

}

/*
@cache
def dfs(i: int, L: int, t: int, e: int) -> int:

	if i == 0:
	    return 1 if L == t == e == 0 else 0
	res = dfs(i - 1, 0, t, e)  # 选 l
	res += dfs(i - 1, L, 0, e)  # 选 t
	res += dfs(i - 1, L, t, max(e - 1, 0))  # 选 e
	res += dfs(i - 1, L, t, e) * 23  # 其它字母
	return res % (10 ** 9 + 7)

class Solution:

	def stringCount(self, n: int) -> int:
	    return dfs(n, 1, 1, 2)
*/
func main() {
	fmt.Println(stringCount(4))
}
