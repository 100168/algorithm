package main

import (
	"fmt"
	"slices"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {

	m := len(nums)
	sort.Ints(nums)
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if nums[i-1]%nums[j-1] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxVal := slices.Max(dp)
	index := 0
	for i := range dp {
		if dp[i] == maxVal {
			index = i
			break
		}
	}

	ans := make([]int, 0)

	pre := nums[index-1]
	ans = append(ans, pre)
	maxVal -= 1
	for maxVal > 0 {
		for j := 1; j < index; j++ {
			if nums[index-1]%nums[j-1] == 0 && dp[j]+1 == dp[index] {
				ans = append(ans, nums[j-1])
				index = j
				maxVal--
				break
			}
		}
	}
	slices.Sort(ans)
	return ans
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{4, 8, 10, 240}))
}
