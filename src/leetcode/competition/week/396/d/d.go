package main

import (
	"math"
	"slices"
)

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {

	mod := int(1e9 + 7)
	s := 0
	n := len(nums)
	maxVal := slices.Max(nums)
	minVal := slices.Min(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != maxVal {
			s += maxVal - nums[i]
		}
	}
	first := s * cost1
	second := math.MaxInt
	up := maxVal * 2

	for i := maxVal; i <= up; i++ {
		d := i - minVal
		cur := 0
		if 2*d-s <= 0 {
			cur = s/2*cost2 + (s%2)*cost1
		} else {
			cur = (s-d)*cost2 + cost1*(2*d-s)
		}
		s += n
		second = min(cur, second)
	}

	return min(first, second) % mod

}

func main() {
	println(minCostToEqualizeArray([]int{1, 14, 14, 15}, 2, 1))
	println(minCostToEqualizeArray([]int{7, 4, 8}, 7, 3))
}
