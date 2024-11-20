package main

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {

	sort.Ints(nums)

	l, r := 0, len(nums)-1

	ans := math.MaxFloat64
	for l < r {
		ans = min(float64(nums[l]+nums[r])/2, ans)
		l++
		r--
	}

	return ans

}
