package main

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {

	n := len(nums)
	ans := make([][]int, 0)
	sort.Ints(nums)

	sMap := make(map[string]bool)
	for i := 0; i < 1<<n; i++ {
		cur := make([]int, 0)
		s := ""
		for j := 0; j < n; j++ {
			if 1<<j&i != 0 {
				cur = append(cur, nums[j])
				s += strconv.Itoa(nums[j])
			}
		}
		if !sMap[s] {
			ans = append(ans, cur)
		}
		sMap[s] = true
	}
	return ans
}

// 000
// 010
// 011
// 100
// 101
// 110
// 111
func subsetsWithDup2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
next:
	for mask := 0; mask < 1<<n; mask++ {
		t := make([]int, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				//去重
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue next
				}
				t = append(t, v)
			}
		}
		ans = append(ans, t)
	}
	return
}
