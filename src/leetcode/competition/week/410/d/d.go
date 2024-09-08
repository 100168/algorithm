package main

import (
	"fmt"
	"slices"
)

/**
给你一个长度为 n 的 正 整数数组 nums 。

如果两个 非负 整数数组 (arr1, arr2) 满足以下条件，我们称它们是 单调 数组对：

两个数组的长度都是 n 。
arr1 是单调 非递减 的，换句话说 arr1[0] <= arr1[1] <= ... <= arr1[n - 1] 。
arr2 是单调 非递增 的，换句话说 arr2[0] >= arr2[1] >= ... >= arr2[n - 1] 。
对于所有的 0 <= i <= n - 1 都有 arr1[i] + arr2[i] == nums[i] 。
请你返回所有 单调 数组对的数目。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。

思路：前缀和优化dp
*/

func countOfPairs(nums []int) int {

	n := len(nums)

	maxVal := slices.Max(nums)

	ans := 0

	mod := int(1e9 + 7)

	f := make([][]int, n)
	s := make([]int, nums[0]+2)

	for i := range f {
		f[i] = make([]int, maxVal+1)
		for j := 0; j <= nums[0]; j++ {
			f[0][j] = 1
			s[j+1] = s[j] + f[0][j]
		}
	}

	for i, v := range nums[1:] {
		s2 := make([]int, v+2)
		for j := 0; j <= v; j++ {

			k := min(nums[i]-v+j, j)
			if k >= 0 {
				f[i+1][j] = s[k+1]
			}
			s2[j+1] = (s2[j] + f[i+1][j]) % mod
		}
		s = s2
	}

	for i := 0; i <= nums[n-1]; i++ {
		ans += f[n-1][i]
	}

	return ans % mod
}

func main() {
	nums := []int{2, 3, 2}
	result := countOfPairs(nums)
	fmt.Println(result) // 输出: 1
}
