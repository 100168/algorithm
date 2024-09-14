package main

import (
	"fmt"
	"slices"
)

/*
*
给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。

如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。

返回 优质数对 的总数。
*/
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {

	nums2Map := make(map[int]int)

	for _, v := range nums2 {
		nums2Map[v]++
	}
	ans := 0
	for _, v := range nums1 {

		if v%k != 0 {
			continue
		}
		v /= k

		for i := 1; i*i <= v; i++ {
			if v%i == 0 {
				ans += nums2Map[i]
				if v/i != i {
					ans += nums2Map[v/i]
				}
			}

		}

	}
	return int64(ans)
}

func numberOfPairs2(nums1, nums2 []int, k int) (ans int64) {
	cnt1 := map[int]int{}
	for _, x := range nums1 {
		if x%k == 0 {
			cnt1[x/k]++
		}
	}
	cnt2 := map[int]int{}
	for _, x := range nums2 {
		cnt2[x]++
	}

	u := slices.Max(nums1) / k
	for i, c := range cnt2 {
		s := 0
		for j := i; j <= u; j += i {
			s += cnt1[j]
		}
		ans += int64(s * c)
	}
	return
}

func main() {
	fmt.Println(numberOfPairs([]int{2, 18, 2, 14}, []int{3, 4, 1, 11}, 2))
	fmt.Println(numberOfPairs2([]int{2, 18, 2, 14}, []int{3, 4, 1, 11}, 2))
}
