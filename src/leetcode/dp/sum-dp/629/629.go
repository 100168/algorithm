package main

import (
	"fmt"
)

/*
*
对于一个整数数组 nums，逆序对是一对满足 0 <= i < j < nums.length 且 nums[i] > nums[j]的整数对 [i, j] 。

给你两个整数 n 和 k，找出所有包含从 1 到 n 的数字，且恰好拥有 k 个 逆序对 的不同的数组的个数。由于答案可能很大，
只需要返回对 109 + 7 取余的结果。

示例 1：

输入：n = 3, k = 0
输出：1
解释：
只有数组 [1,2,3] 包含了从1到3的整数并且正好拥有 0 个逆序对。
示例 2：

输入：n = 3, k = 1
输出：2
解释：
数组 [1,3,2] 和 [2,1,3] 都有 1 个逆序对。

提示：

1 <= n <= 1000
0 <= k <= 1000

思路：前缀和优化dp

思路：需要沿途累加，需要注意取值范围
*/

const mod = 1_000_000_007

func kInversePairs(n int, k int) int {

	f := make([]int, k+1)
	s := make([]int, k+2)
	for i := range s {
		s[i] = 1
	}
	s[0] = 0
	for i := 1; i <= n; i++ {
		t := make([]int, k+2)
		for j := 0; j <= k; j++ {
			f[j] = (s[j+1] - s[max(j-i+1, 0)] + mod) % mod
			t[j+1] = (t[j] + f[j]) % mod
		}
		s = t
	}
	return f[k]
}
func main() {
	fmt.Println(kInversePairs(3, 1))
}