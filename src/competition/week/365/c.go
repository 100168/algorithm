package main

import "math"

func minSizeSubarray(nums []int, target int) int {

	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	rest := target % sum
	cnt := target / sum
	if rest == 0 {
		return n * cnt
	}
	ans := math.MaxInt
	left := 0
	curSum := 0
	for i := 0; i < 2*n; i++ {
		curSum += nums[i%n]
		for curSum > rest && left < i {
			curSum -= nums[left%n]
			left++
		}
		if curSum == rest {
			if i-left+1 < ans {
				ans = i - left + 1
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + cnt*n
}
func main() {
	println(minSizeSubarray([]int{1, 2}, 78))
}
