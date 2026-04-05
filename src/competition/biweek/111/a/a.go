package main

import "sort"

func countPairs(nums []int, target int) int {

	sort.Ints(nums)
	left := 0
	right := len(nums) - 1

	ans := 0
	for left < right {
		cur := nums[left] + nums[right]
		if cur < target {
			ans += right - left
			left++
		} else if cur >= target {
			right--
		}

	}
	return ans

}
func main() {

	nums := []int{-1, 1, 2, 3, 1}
	println(countPairs(nums, 2))
}
