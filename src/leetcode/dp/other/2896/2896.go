package main

import (
	"fmt"
	"math"
)

/*
*
给你两个下标从 0 开始的二进制字符串 s1 和 s2 ，两个字符串的长度都是 n ，再给你一个正整数 x 。

你可以对字符串 s1 执行以下操作 任意次 ：

选择两个下标 i 和 j ，将 s1[i] 和 s1[j] 都反转，操作的代价为 x 。
选择满足 i < n - 1 的下标 i ，反转 s1[i] 和 s1[i + 1] ，操作的代价为 1 。
请你返回使字符串 s1 和 s2 相等的 最小 操作代价之和，如果无法让二者相等，返回 -1 。

注意 ，反转字符的意思是将 0 变成 1 ，或者 1 变成 0 。

示例 1：

输入：s1 = "1100011000", s2 = "0101001010", x = 2
输出：4
解释：我们可以执行以下操作：
- 选择 i = 3 执行第二个操作。结果字符串是 s1 = "1101111000" 。
- 选择 i = 4 执行第二个操作。结果字符串是 s1 = "1101001000" 。
- 选择 i = 0 和 j = 8 ，执行第一个操作。结果字符串是 s1 = "0101001010" = s2 。
总代价是 1 + 1 + 2 = 4 。这是最小代价和。
*/

func minOperations(s1 string, s2 string, x int) int {

	n := len(s1)

	diff := make([]bool, n)

	xor := uint8(0)
	for i := 0; i < n; i++ {
		xor ^= s1[i] ^ s2[i]
		diff[i] = s1[i] != s2[i]
	}

	if xor > 0 {
		return -1
	}

	f := make(map[int]int)

	var dfs func(int, int, bool) int

	dfs = func(i int, ops int, pre bool) int {

		if i < 0 {

			if ops == 0 && !pre {
				return 0
			}
			return math.MaxInt / 2
		}

		key := i<<11 + ops<<1
		if pre {
			key += 1
		}
		if _, ok := f[key]; ok {
			return f[key]
		}

		cur := math.MaxInt / 2

		if pre == diff[i] {
			cur = min(cur, dfs(i-1, ops, false))
		} else {
			cur = dfs(i-1, ops+1, false) + x
			cur = min(cur, dfs(i-1, ops, true)+1)
			if ops > 0 {
				cur = min(cur, dfs(i-1, ops-1, false))
			}
		}
		f[key] = cur
		return cur

	}
	ans := dfs(n-1, 0, false)

	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}

func minOperations2(s1 string, s2 string, x int) int {

	n := len(s1)

	diff := make([]bool, n)

	cur := uint8(0)
	for i := 0; i < n; i++ {
		cur ^= s1[i] ^ s2[i]
		diff[i] = s1[i] != s2[i]
	}

	if cur > 0 {
		return -1
	}

	memo := make([][]map[bool]int, n)

	for i := range memo {
		memo[i] = make([]map[bool]int, n)

		for j := range memo[i] {
			memo[i][j] = make(map[bool]int)
		}
	}

	var dfs func(int, int, bool) int

	dfs = func(i, j int, pre bool) int {
		if i < 0 {
			if pre || j > 0 {
				return math.MaxInt / 2
			}
			return 0
		}

		if _, ok := memo[i][j][pre]; ok {
			return memo[i][j][pre]
		}
		cur := math.MaxInt / 2

		if pre == diff[i] {
			cur = dfs(i-1, j, false)
		} else {
			cur = dfs(i-1, j+1, false) + x
			cur = min(cur, dfs(i-1, j, true)+1)
			if j > 0 {
				cur = min(cur, dfs(i-1, j-1, false))
			}
		}
		memo[i][j][pre] = cur
		return cur
	}

	return dfs(n-1, 0, false)

}

func minOperations3(s1 string, s2 string, x int) int {

	n := len(s1)

	f := make(map[int]int)

	var dfs func(int, int, bool) int

	dfs = func(i int, ops int, pre bool) int {

		if i < 0 {

			if ops == 0 && !pre {
				return 0
			}
			return math.MaxInt / 2
		}

		key := i<<11 + ops<<1
		if pre {
			key += 1
		}
		if _, ok := f[key]; ok {
			return f[key]
		}

		cur := math.MaxInt / 2

		if pre {
			if s1[i] != s2[i] {
				cur = min(cur, dfs(i-1, ops, false))
			} else {
				cur = min(cur, dfs(i-1, ops, true)+1)
				cur = min(cur, dfs(i-1, ops+1, false)+x)
			}
		} else {
			if s1[i] == s2[i] {
				cur = min(cur, dfs(i-1, ops, false))
			} else {
				if ops > 0 {
					cur = min(cur, dfs(i-1, ops-1, false))
				} else {
					cur = min(cur, dfs(i-1, ops+1, false)+x)
				}
				cur = min(cur, dfs(i-1, ops, true)+1)
			}
		}
		f[key] = cur

		return cur

	}
	ans := dfs(n-1, 0, false)

	if ans == math.MaxInt/2 {
		return -1
	}
	return ans
}
func main() {
	fmt.Println(minOperations("1100011000", "0101001010", 2))
}
