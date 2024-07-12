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
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = abs(nums1[i] - nums2[i])
	}
	t := k1 + k2
	sort.Ints(diff)

	l, r := 0, diff[n-1]

	check := func(m int) bool {

		s := 0
		for i := n - 1; i >= 0; i-- {
			if diff[i] <= m {
				break
			}
			s += diff[i] - m

		}
		return t >= s
	}

	for l <= r {
		m := l + (r-l)/2
		if check(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if diff[i] <= l {
			break
		}
		t -= diff[i] - l
		diff[i] = l
	}
	if diff[n-1] == 0 {
		return 0

	}
	for i := n - 1; i >= 0 && t > 0; i-- {
		diff[i]--
		t--
	}
	ans := 0

	for _, v := range diff {
		ans += v * v
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
	fmt.Println(minSumSquareDiff([]int{10, 10, 10, 11, 5}, []int{1, 0, 6, 6, 1}, 11, 27))
}
