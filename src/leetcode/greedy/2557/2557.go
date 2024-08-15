package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*
给你一个整数数组 banned 和两个整数 n 和 maxSum 。你需要按照以下规则选择一些整数：

被选择整数的范围是 [1, n] 。
每个整数 至多 选择 一次 。
被选择整数不能在数组 banned 中。
被选择整数的和不超过 maxSum 。
请你返回按照上述规则 最多 可以选择的整数数目。

示例 1：

输入：banned = [1,4,6], n = 6, maxSum = 4
输出：1
解释：你可以选择整数 3 。
3 在范围 [1, 6] 内，且不在 banned 中，所选整数的和为 3 ，也没有超过 maxSum 。
*/
func maxCount(banned []int, n int, maxSum int64) int {

	sort.Ints(banned)
	banned = slices.Compact(banned)
	sum := make([]int, len(banned)+1)
	for i, v := range banned {
		sum[i+1] = sum[i] + v
	}
	l, r := 1, n

	check := func(t int) bool {
		l1, r1 := 0, len(banned)-1

		for l1 <= r1 {
			m := (l1 + r1) / 2
			if banned[m] > t {
				r1 = m - 1
			} else {
				l1 = m + 1
			}
		}
		s1 := sum[l1]

		s2 := (1 + t) * t / 2
		return int64(s2-s1) <= maxSum
	}
	for l <= r {
		m := (l + r) / 2
		if check(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	l1, r1 := 0, len(banned)-1

	for l1 <= r1 {
		m := (l1 + r1) / 2
		if banned[m] > l {
			r1 = m - 1
		} else {
			l1 = m + 1
		}
	}

	return r - l1
}

func main() {
	//fmt.Println(maxCount([]int{1, 4, 6}, 1<<15, 1<<19))
	fmt.Println(maxCount([]int{1, 1}, 2, 2))
}
