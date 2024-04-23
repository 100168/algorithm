package main

import (
	"math"
	"slices"
	"sort"
)

func longestSquareStreak(nums []int) int {

	numsMap := make(map[int]int)

	for _, v := range nums {

		numsMap[v] = 1
	}

	sort.Ints(nums)

	nums = slices.Compact(nums)

	ans := -1
	for _, v := range nums {
		ads := int(math.Sqrt(float64(v)))
		if ads*ads == v {
			if _, ok := numsMap[ads]; ok {
				numsMap[v] += numsMap[ads]
				ans = max(ans, numsMap[v])
			}
		}
	}
	return ans
}

func main() {

	println(longestSquareStreak([]int{2, 4, 4, 5, 5}))
}
