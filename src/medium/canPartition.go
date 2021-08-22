package main

import "fmt"

//0-1 bag
func canPartition(nums []int) bool {

	if len(nums) <= 1 {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	if sum%2 == 1 {
		return false
	}

	midSum := sum / 2
	if max > midSum {
		return false
	}
	return process2(nums, 0, midSum)
}

func process2(nums []int, index int, value int) bool {
	m := len(nums)
	dp := make([][]bool, m)

	for i := range dp {
		dp[i] = make([]bool, value+1)
	}
	dp[0][nums[0]] = true
	for i := 0; i < m; i++ {
		dp[i][0] = true
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= value; j++ {
			diff := j - nums[i]
			if diff < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][diff]
			}
		}
	}
	return dp[m-1][value]
}
func process(nums []int, index int, value int) bool {
	if value == 0 {
		return true
	}
	if index == len(nums) && value != 0 {
		return false
	}

	ans := false
	ans = ans || process(nums, index+1, value)
	if ans {
		return ans
	}
	ans = ans || process(nums, index+1, value-nums[index])
	return ans
}

func main() {
	nums := []int{1, 5, 11, 5}
	partition := canPartition(nums)
	fmt.Println(partition)
}
