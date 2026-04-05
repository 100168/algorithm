package main

import "math"

func findMaxAverage(nums []int, k int) float64 {

	left := math.MaxFloat64
	right := float64(0)

	for _, v := range nums {
		if float64(v) < left {
			left = float64(v)
		}
		if float64(v) > right {
			right = float64(v)
		}
	}

	for right-left > 0.00001 {
		mid := (left + right) * 0.5

		if check(mid, nums, k) {
			left = mid
		} else {
			right = mid
		}

	}

	return left

}

func check(mid float64, nums []int, k int) bool {
	sum := float64(0)
	pre := float64(0)
	minSum := math.MaxFloat64

	for i := 0; i < k; i++ {
		sum += float64(nums[i]) - mid
	}

	if sum >= float64(0) {
		return true
	}

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i]) - mid
		pre += float64(nums[i-k]) - mid

		if pre < minSum {
			minSum = pre
		}

		if sum-minSum >= float64(0) {
			return true
		}
	}

	return false
}
