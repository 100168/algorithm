package main

import (
	"fmt"
	"sort"
)

/**

给你一个 二进制 字符串 s 和一个整数 k。

另给你一个二维整数数组 queries ，其中 queries[i] = [li, ri] 。

如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：

字符串中 0 的数量最多为 k。
字符串中 1 的数量最多为 k。
返回一个整数数组 answer ，其中 answer[i] 表示 s[li..ri] 中满足 k 约束 的
子字符串
 的数量。

示例 1：

输入：s = "0001111", k = 2, queries = [[0,6]]

输出：[26]

解释：

对于查询 [0, 6]， s[0..6] = "0001111" 的所有子字符串中，除 s[0..5] = "000111" 和 s[0..6] = "0001111" 外，其余子字符串都满足 k 约束。
*/

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	n := len(s)

	//以i为右端点，最远左端点是多少
	left := make([]int, n)
	sum := make([]int, n+1)
	cnt := [2]int{}
	l := 0
	for i, c := range s {
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]&1]--
			l++
		}
		left[i] = l
		// 计算 i-left[i]+1 的前缀和
		sum[i+1] = sum[i] + i - l + 1
	}

	//  l  r
	ans := make([]int64, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		j := l + sort.SearchInts(left[l:r+1], l)
		ans[i] = int64(sum[r+1] - sum[j] + (j-l+1)*(j-l)/2)
	}
	return ans
}

func main() {
	fmt.Println(countKConstraintSubstrings("0001111", 2, [][]int{{0, 6}}))
}
