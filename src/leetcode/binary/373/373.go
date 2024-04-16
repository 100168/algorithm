package main

import (
	"sort"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	m := len(nums1)
	n := len(nums2)
	l, r := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]

	for l <= r {
		mid := (r + l) / 2
		if check(mid, nums1, nums2) <= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	i := 0
	j := n - 1
	ans := make([][]int, 0)
	for i < m && j >= 0 {
		if nums1[i]+nums2[j] >= r {
			j--
		} else {
			for c := 0; c <= j; c++ {
				cur := []int{nums1[i], nums2[c]}
				ans = append(ans, cur)
			}
			i++
		}
	}

	if len(ans) == k {
		return ans
	}

	i = 0
	j = n - 1
	for i < m && j >= 0 {
		if nums1[i]+nums2[j] > r {
			j--
		} else {
			for c := 0; c <= j; c++ {

				if nums1[i]+nums2[c] == r {
					cur := []int{nums1[i], nums2[c]}
					ans = append(ans, cur)
				}
				if len(ans) == k {
					return ans
				}
			}
			i++
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0]+ans[i][1] == ans[j][0]+ans[j][1] {
			return ans[i][0] < ans[j][0]
		} else {
			return ans[i][0]+ans[i][1] < ans[j][0]+ans[j][1]
		}
	})
	return ans
}

func check(t int, nums1, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)

	i := 0
	j := n - 1

	cnt := 1
	for i < m && j >= 0 {

		if nums1[i]+nums2[j] >= t {
			j--
		} else {
			i++
			cnt += j + 1
		}
	}
	return cnt
}

func main() {

	//println(kSmallestPairs([]int{2, 4, 6}, []int{1, 7, 11}, 3))
	println(kSmallestPairs([]int{1, 1, 1}, []int{1, 1, 1}, 5))
}
