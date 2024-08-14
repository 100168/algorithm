package main

import (
	"fmt"
	"sort"
)

/**
给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，长度为 n 。

数组 nums1 和 nums2 的 差值平方和 定义为所有满足 0 <= i < n 的 (nums1[i] - nums2[i])2 之和。

同时给你两个正整数 k1 和 k2 。你可以将 nums1 中的任意元素 +1 或者 -1 至多 k1 次。类似的，你可以将 nums2 中的任意元素 +1 或者 -1 至多 k2 次。

请你返回修改数组 nums1 至多 k1 次且修改数组 nums2 至多 k2 次后的最小 差值平方和 。

注意：你可以将数组中的元素变成 负 整数。
*/

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {

	n := len(nums1)
	a := make([]int, n+1)

	ans := 0
	s := 0
	for i := 0; i < n; i++ {
		a[i+1] = abs(nums1[i] - nums2[i])
		ans += abs(nums1[i]-nums2[i]) * abs(nums1[i]-nums2[i])
		s += abs(nums1[i] - nums2[i])
	}

	k := k1 + k2
	if k >= s {
		return 0
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	for i, v := range a {
		ans -= v * v
		i++
		if k >= i*(v-a[i]) {
			k -= i * (v - a[i])
			continue
		}

		v -= k / i

		ans += (v-1)*(v-1)*(k%i) + v*v*(i-k%i)
		break

	}
	return int64(ans)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//fmt.Println(minSumSquareDiff([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	//fmt.Println(minSumSquareDiff2([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
	//fmt.Println(minSumSquareDiff([]int{10, 10, 10, 11, 5}, []int{1, 0, 6, 6, 1}, 11, 27))
	fmt.Println(minSumSquareDiff([]int{11, 12, 13, 14, 15}, []int{13, 16, 16, 12, 14}, 3, 6))
}
