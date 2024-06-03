package main

import (
	"fmt"
	"math"
)

/*
*给你一个正整数数组 nums，请你移除 最短 子数组（可以为 空），使得剩余元素的 和 能被 p 整除。 不允许 将整个数组都移除。

请你返回你需要移除的最短子数组的长度，如果无法满足题目要求，返回 -1 。

子数组 定义为原数组中连续的一组元素。
*/
func minSubarray(nums []int, p int) int {

	sMap := make(map[int]int)
	suffixes := make([]int, len(nums)+1)

	ans := math.MaxInt
	s := 0
	for i := len(nums) - 1; i >= 0; i-- {
		s = (s + nums[i]) % p
		if s == 0 {
			ans = min(ans, i)
		}
		suffixes[i] = s
	}

	s = 0

	if suffixes[0] == 0 {
		return 0
	}
	for i, v := range nums {
		s = (s + v) % p

		if s == 0 {
			ans = min(ans, len(nums)-i-1)
		}
		suffix := suffixes[i+1]
		//((s+p)-suffix)%p
		if _, ok := sMap[(p-suffix)%p]; ok {
			ans = min(ans, i-sMap[(p-suffix)%p])
		}
		sMap[s] = i
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minSubarray2(nums []int, p int) int {

	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	remainded := sum % p
	if remainded == 0 {

		return remainded
	}

	exist := map[int]int{0: -1}
	cur := 0

	ans := len(nums)
	for i, num := range nums {
		cur = (cur + num) % p
		key := ((cur - remainded) + p) % p
		if _, ok := exist[key]; ok {
			ans = min(ans, i-exist[key])
		}
		exist[cur] = i
	}
	if ans >= n {
		return -1
	}
	return ans

}

func main() {
	fmt.Println(minSubarray([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148))
}
