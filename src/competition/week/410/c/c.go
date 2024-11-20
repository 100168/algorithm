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
*/

func countOfPairs(nums []int) int {

	mod := int(1e9 + 7)

	n := len(nums)

	maxV := slices.Max(nums)

	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, maxV+1)
		for j := range f[i] {
			f[i][j] = make([]int, maxV+1)
			for k := range f[i][j] {
				f[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int

	dfs = func(i int, first int, second int) int {

		if i < 0 {
			return 1
		}

		if f[i][first][second] != -1 {
			return f[i][first][second]
		}
		cur := 0
		v := nums[i]
		for j := 0; j <= v; j++ {

			if j <= first && (v-j) >= second {
				second2 := v - j
				cur += dfs(i-1, j, second2)
				cur %= mod
			}
		}
		f[i][first][second] = cur
		return cur

	}

	ans := 0
	for i := 0; i <= nums[n-1]; i++ {

		ans += dfs(n-2, i, nums[n-1]-i)
		ans %= mod

	}
	return ans

}

func main() {
	fmt.Println(countOfPairs([]int{2, 3, 2}))
}
