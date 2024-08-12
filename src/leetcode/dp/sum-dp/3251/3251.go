package main

import (
	"fmt"
	"slices"
)

/*
*
给你一个长度为 n 的 正 整数数组 nums 。

如果两个 非负 整数数组 (arr1, arr2) 满足以下条件，我们称它们是 单调 数组对：

两个数组的长度都是 n 。
arr1 是单调 非递减 的，换句话说 arr1[0] <= arr1[1] <= ... <= arr1[n - 1] 。
arr2 是单调 非递增 的，换句话说 arr2[0] >= arr2[1] >= ... >= arr2[n - 1] 。
对于所有的 0 <= i <= n - 1 都有 arr1[i] + arr2[i] == nums[i] 。
请你返回所有 单调 数组对的数目。

由于答案可能很大，请你将它对 109 + 7 取余 后返回。

思路：
前缀和优化dp
f[i]
f[v-i],
*/
func countOfPairs(nums []int) int {

	mod := int(1e9 + 7)
	n := len(nums)
	count := 0

	maxVal := slices.Max(nums)
	f1 := make([]int, maxVal+1)
	f2 := make([]int, maxVal+1)

	for i := 0; i <= nums[0]; i++ {
		if i == 0 {
			f2[i] = 1
			f1[nums[i]-i] = +1

		} else {
			f2[i] = f2[i-1] + 1
			f1[nums[i]-i] = f1[nums[i]-i+1] + 1
		}
	}
	//k<=j =>nums[i-1]-k>=nums[i]-j=>   k<=min(nums[i-1]-nums[i]+j,j) ==>k<= min(nums[i-1]-nums[i],0)+j

	for i := 1; i < n; i++ {
		f11 := make([]int, maxVal+1)
		f22 := make([]int, maxVal+1)
		c := nums[i]
		pre := nums[i-1]
		for j := 0; j <= c; j++ {
			less := 0
			more := f2[c-j]
			if j > pre {
				less = f1[0]
			} else {
				less = f1[j]
			}
			if more != 0 {
				f11[j] = min(less, more) + 1
				f22[c-j] = min(less, more) + 1
			}
		}
		f1 = f11
		f2 = f22

	}

	for _, v := range f1 {
		count += v
		v %= mod
	}

	return count

}

func countOfPairs2(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	s := make([]int, m+1)

	for j := 0; j <= nums[0]; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		s[0] = f[i-1][0]
		for k := 1; k <= m; k++ {
			s[k] = s[k-1] + f[i-1][k] // f[i-1] 的前缀和
		}
		for j := 0; j <= nums[i]; j++ {
			maxK := j + min(nums[i-1]-nums[i], 0)
			if maxK >= 0 {
				f[i][j] = s[maxK] % mod
			}
		}
	}

	for _, v := range f[n-1][:nums[n-1]+1] {
		ans += v
	}
	return ans % mod
}

func main() {
	fmt.Println(countOfPairs([]int{2, 3, 2}))
}
