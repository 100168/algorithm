package main

import "math"

func maxScore(nums []int, x int) int64 {

	n := len(nums)
	odd := int64(math.MinInt / 2)
	even := int64(math.MinInt / 2)

	if nums[0]%2 == 0 {
		even = int64(nums[0])
	} else {
		odd = int64(nums[0])
	}
	for i := 1; i < n; i++ {
		v := nums[i]
		if v&1 == 0 {
			even = max(even+int64(v), odd+int64(v)-int64(x))
		} else {
			odd = max(odd+int64(v), even+int64(v)-int64(x))
		}

	}
	return max(odd, even)
}
