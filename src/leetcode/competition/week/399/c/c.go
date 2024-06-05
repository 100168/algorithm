package main

import "fmt"

/*
*
给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。

如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。

返回 优质数对 的总数。
*/
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {

	nums2Map := make(map[int]int)

	for _, n := range nums2 {
		nums2Map[n]++
	}

	ans := 0

	for _, n := range nums1 {
		if n%k == 0 {
			cur := n / k
			for j := 1; j*j <= cur; j++ {

				if cur%j == 0 {
					ans += nums2Map[j]
					if j != cur/j {
						ans += nums2Map[cur/j]
					}
				}

			}

		}
	}
	return int64(ans)
}

func numberOfPairs2(nums1 []int, nums2 []int, k int) int {

	ans := 0
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i]%(nums2[j]*k) == 0 {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfPairs([]int{2, 18, 2, 14}, []int{3, 4, 1, 11}, 2))
	fmt.Println(numberOfPairs2([]int{2, 18, 2, 14}, []int{3, 4, 1, 11}, 2))
}
